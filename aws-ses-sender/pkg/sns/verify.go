// Package sns verifies the messages that Amazon SNS posts to an HTTPS endpoint.
//
// Anyone can post JSON to a public URL, so an endpoint that receives SNS messages must check
// that AWS really sent them. SNS signs every message with a certificate that it publishes, and
// the message names that certificate. The signature alone is not enough, because anybody with an
// AWS account can create a topic of their own and have SNS sign messages for it. A Verifier
// therefore also requires the message to come from one of the topics it was created for.
//
// See "Verifying the signatures of Amazon SNS messages" in the Amazon SNS developer guide.
package sns

import (
	"context"
	"crypto"
	"crypto/rsa"
	_ "crypto/sha1"   // registers SHA-1, which SignatureVersion 1 of SNS signs with
	_ "crypto/sha256" // registers SHA-256, which SignatureVersion 2 signs with
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"
)

// The message types SNS sends to an HTTPS endpoint.
const (
	TypeNotification             = "Notification"
	TypeSubscriptionConfirmation = "SubscriptionConfirmation"
	TypeUnsubscribeConfirmation  = "UnsubscribeConfirmation"
)

const (
	// maxCertSize bounds what is read from a certificate URL, a real one is about 2 KB.
	maxCertSize = 64 << 10
	// maxCachedCerts bounds the certificate cache. Only the few certificates AWS publishes can
	// get into it, because the URL has to pass checkCertURL first.
	maxCachedCerts = 32
)

// certHost matches the hosts that serve SNS signing certificates, for example
// sns.us-east-1.amazonaws.com and sns.cn-north-1.amazonaws.com.cn.
var certHost = regexp.MustCompile(`^sns\.[a-z0-9-]+\.amazonaws\.com(\.cn)?$`)

// Message is the JSON document SNS posts. Fields that a message type does not have stay empty.
type Message struct {
	Type             string `json:"Type"`
	MessageID        string `json:"MessageId"`
	Token            string `json:"Token"`
	TopicARN         string `json:"TopicArn"`
	Subject          string `json:"Subject"`
	Message          string `json:"Message"`
	SubscribeURL     string `json:"SubscribeURL"`
	Timestamp        string `json:"Timestamp"`
	SignatureVersion string `json:"SignatureVersion"`
	Signature        string `json:"Signature"`
	SigningCertURL   string `json:"SigningCertURL"`
}

// stringToSign builds the text that SNS signed: the name and the value of each field the
// message type signs, one per line, in alphabetical order. A Notification without a subject
// leaves the subject out.
func (m *Message) stringToSign() (string, error) {
	var names []string
	switch m.Type {
	case TypeNotification:
		names = []string{"Message", "MessageId", "Subject", "Timestamp", "TopicArn", "Type"}
	case TypeSubscriptionConfirmation, TypeUnsubscribeConfirmation:
		names = []string{"Message", "MessageId", "SubscribeURL", "Timestamp", "Token", "TopicArn", "Type"}
	default:
		return "", fmt.Errorf("sns: unsupported message type %q", m.Type)
	}

	values := map[string]string{
		"Message":      m.Message,
		"MessageId":    m.MessageID,
		"Subject":      m.Subject,
		"SubscribeURL": m.SubscribeURL,
		"Timestamp":    m.Timestamp,
		"Token":        m.Token,
		"TopicArn":     m.TopicARN,
		"Type":         m.Type,
	}

	var b strings.Builder
	for _, name := range names {
		if name == "Subject" && m.Subject == "" {
			continue
		}
		b.WriteString(name)
		b.WriteByte('\n')
		b.WriteString(values[name])
		b.WriteByte('\n')
	}
	return b.String(), nil
}

// Verifier verifies SNS messages. It is safe for concurrent use.
type Verifier struct {
	// Client downloads the signing certificates. NewVerifier sets one with a timeout that does
	// not follow redirects.
	Client *http.Client

	topics map[string]struct{}

	mu    sync.Mutex
	certs map[string]*rsa.PublicKey
}

// NewVerifier returns a Verifier that accepts messages from the given topic ARNs only. Without
// any topic it rejects every message, so an endpoint that was not configured fails closed.
func NewVerifier(topicARNs ...string) *Verifier {
	topics := make(map[string]struct{}, len(topicARNs))
	for _, arn := range topicARNs {
		if arn != "" {
			topics[arn] = struct{}{}
		}
	}
	return &Verifier{
		Client: &http.Client{
			Timeout: 10 * time.Second,
			// A certificate URL that redirects is not one that AWS published.
			CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse },
		},
		topics: topics,
		certs:  make(map[string]*rsa.PublicKey),
	}
}

// Configured reports whether the Verifier accepts messages from at least one topic.
func (v *Verifier) Configured() bool {
	return len(v.topics) > 0
}

// Verify returns nil only if m comes from one of the Verifier's topics and carries a valid
// signature by the certificate that it names.
func (v *Verifier) Verify(ctx context.Context, m *Message) error {
	if len(v.topics) == 0 {
		return errors.New("sns: no topic is configured")
	}
	if _, ok := v.topics[m.TopicARN]; !ok {
		return fmt.Errorf("sns: message is from topic %q, which is not configured", m.TopicARN)
	}

	var hash crypto.Hash
	switch m.SignatureVersion {
	case "1":
		hash = crypto.SHA1
	case "2":
		hash = crypto.SHA256
	default:
		return fmt.Errorf("sns: unsupported signature version %q", m.SignatureVersion)
	}

	signed, err := m.stringToSign()
	if err != nil {
		return err
	}
	signature, err := base64.StdEncoding.DecodeString(m.Signature)
	if err != nil || len(signature) == 0 {
		// Refuse before downloading a certificate: an unsigned message is not worth the request.
		return errors.New("sns: signature is missing or is not base64")
	}

	key, err := v.publicKey(ctx, m.SigningCertURL)
	if err != nil {
		return err
	}

	h := hash.New()
	h.Write([]byte(signed))
	if err := rsa.VerifyPKCS1v15(key, hash, h.Sum(nil), signature); err != nil {
		return fmt.Errorf("sns: signature does not match: %w", err)
	}
	return nil
}

// checkCertURL accepts only the https URLs of .pem files on the SNS hosts of AWS.
func checkCertURL(raw string) (*url.URL, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return nil, fmt.Errorf("sns: invalid signing certificate URL: %w", err)
	}
	switch {
	case u.Scheme != "https":
		return nil, fmt.Errorf("sns: signing certificate URL %q is not https", raw)
	case u.User != nil:
		return nil, fmt.Errorf("sns: signing certificate URL %q has credentials", raw)
	case u.Port() != "" && u.Port() != "443":
		return nil, fmt.Errorf("sns: signing certificate URL %q uses another port", raw)
	case !certHost.MatchString(strings.ToLower(u.Hostname())):
		return nil, fmt.Errorf("sns: signing certificate URL %q is not an SNS host", raw)
	case !strings.HasSuffix(u.Path, ".pem") || u.RawQuery != "" || u.Fragment != "":
		return nil, fmt.Errorf("sns: signing certificate URL %q is not a .pem file", raw)
	}
	return u, nil
}

// publicKey returns the public key of the certificate at raw, from the cache if it is there.
func (v *Verifier) publicKey(ctx context.Context, raw string) (*rsa.PublicKey, error) {
	u, err := checkCertURL(raw)
	if err != nil {
		return nil, err
	}

	v.mu.Lock()
	key, ok := v.certs[u.String()]
	v.mu.Unlock()
	if ok {
		return key, nil
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := v.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sns: downloading the signing certificate: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sns: downloading the signing certificate: status %s", resp.Status)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, maxCertSize))
	if err != nil {
		return nil, fmt.Errorf("sns: reading the signing certificate: %w", err)
	}

	block, _ := pem.Decode(body)
	if block == nil || block.Type != "CERTIFICATE" {
		return nil, errors.New("sns: the signing certificate is not PEM encoded")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("sns: parsing the signing certificate: %w", err)
	}
	key, ok = cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("sns: the signing certificate does not hold an RSA key")
	}

	v.mu.Lock()
	if len(v.certs) < maxCachedCerts {
		v.certs[u.String()] = key
	}
	v.mu.Unlock()
	return key, nil
}
