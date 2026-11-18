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
	"encoding/json"
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
	// maxCachedCerts bounds the certificate cache. Only a certificate that has verified the
	// signature of a message from a configured topic is cached, so a caller cannot fill it.
	maxCachedCerts = 32
)

// certHost matches the hosts that serve SNS signing certificates: sns.<region>.amazonaws.com,
// for example sns.us-east-1.amazonaws.com, sns.us-gov-west-1.amazonaws.com and
// sns.cn-north-1.amazonaws.com.cn. The region has the shape of a region, which leaves out
// names such as sns.s3.amazonaws.com.
var certHost = regexp.MustCompile(`^sns\.[a-z]{2}(-[a-z]+)+-[0-9]+\.amazonaws\.com(\.cn)?$`)

// ErrMalformed is returned (wrapped) by Receive for a request that is not a well-formed SNS
// message of a supported type. A message that is well formed but cannot be verified gives
// another error.
var ErrMalformed = errors.New("sns: malformed message")

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
		return "", fmt.Errorf("sns: unsupported message type %.100q", m.Type)
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

// Receive reads the body of a request that SNS sent to the endpoint and returns the message
// only if Verify accepts it, so a handler cannot use a message that was not verified.
// messageType is the value of the x-amz-sns-message-type header. It has to be Notification or
// SubscriptionConfirmation, and the body has to say the same. Anything else is ErrMalformed.
func (v *Verifier) Receive(ctx context.Context, messageType string, body []byte) (*Message, error) {
	if messageType != TypeNotification && messageType != TypeSubscriptionConfirmation {
		return nil, fmt.Errorf("%w: unsupported message type %.100q", ErrMalformed, messageType)
	}
	var m Message
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrMalformed, err)
	}
	if m.Type != messageType {
		return nil, fmt.Errorf("%w: the header says %q and the body says %.100q", ErrMalformed, messageType, m.Type)
	}
	if err := v.Verify(ctx, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

// Verify returns nil only if m comes from one of the Verifier's topics and carries a valid
// signature by the certificate that it names.
func (v *Verifier) Verify(ctx context.Context, m *Message) error {
	if len(v.topics) == 0 {
		return errors.New("sns: no topic is configured")
	}
	if _, ok := v.topics[m.TopicARN]; !ok {
		return fmt.Errorf("sns: message is from topic %.100q, which is not configured", m.TopicARN)
	}

	var hash crypto.Hash
	switch m.SignatureVersion {
	case "1":
		hash = crypto.SHA1
	case "2":
		hash = crypto.SHA256
	default:
		return fmt.Errorf("sns: unsupported signature version %.100q", m.SignatureVersion)
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

	certURL, err := checkCertURL(m.SigningCertURL)
	if err != nil {
		return err
	}
	key, cached := v.cachedKey(certURL)
	if !cached {
		if key, err = v.download(ctx, certURL); err != nil {
			return err
		}
	}

	h := hash.New()
	h.Write([]byte(signed))
	if err := rsa.VerifyPKCS1v15(key, hash, h.Sum(nil), signature); err != nil {
		return fmt.Errorf("sns: signature does not match: %w", err)
	}

	// Only now is the certificate known to have signed a message of a configured topic, so only
	// now is it cached.
	if !cached {
		v.remember(certURL, key)
	}
	return nil
}

// checkCertURL accepts only the https URLs of .pem files on the SNS hosts of AWS and returns
// them in one canonical spelling (lower-case host, no port), which is also the cache key.
func checkCertURL(raw string) (string, error) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", fmt.Errorf("sns: invalid signing certificate URL: %w", err)
	}
	host := strings.ToLower(u.Hostname())
	switch {
	case u.Scheme != "https":
		return "", fmt.Errorf("sns: signing certificate URL %.100q is not https", raw)
	case u.User != nil:
		return "", fmt.Errorf("sns: signing certificate URL %.100q has credentials", raw)
	case u.Port() != "" && u.Port() != "443":
		return "", fmt.Errorf("sns: signing certificate URL %.100q uses another port", raw)
	case !certHost.MatchString(host):
		return "", fmt.Errorf("sns: signing certificate URL %.100q is not an SNS host", raw)
	case !strings.HasSuffix(u.Path, ".pem") || strings.Contains(u.EscapedPath(), "%") || u.RawQuery != "" || u.Fragment != "":
		return "", fmt.Errorf("sns: signing certificate URL %.100q is not a plain .pem file", raw)
	}
	return "https://" + host + u.EscapedPath(), nil
}

// cachedKey returns the key of a certificate that has been used before.
func (v *Verifier) cachedKey(certURL string) (*rsa.PublicKey, bool) {
	v.mu.Lock()
	defer v.mu.Unlock()
	key, ok := v.certs[certURL]
	return key, ok
}

// remember caches a key. A full cache stays as it is.
func (v *Verifier) remember(certURL string, key *rsa.PublicKey) {
	v.mu.Lock()
	defer v.mu.Unlock()
	if len(v.certs) < maxCachedCerts {
		v.certs[certURL] = key
	}
}

// download fetches the certificate and returns its public key.
func (v *Verifier) download(ctx context.Context, certURL string) (*rsa.PublicKey, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, certURL, nil)
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
	key, ok := cert.PublicKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("sns: the signing certificate does not hold an RSA key")
	}
	return key, nil
}
