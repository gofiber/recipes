package sns

import (
	"bytes"
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"io"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	testTopic   = "arn:aws:sns:ap-northeast-2:123456789012:ses-events"
	testCertURL = "https://sns.ap-northeast-2.amazonaws.com/SimpleNotificationService-0123456789abcdef0123456789abcdef.pem"
)

var (
	keyOnce sync.Once
	testKey *rsa.PrivateKey
)

// signingKey returns one RSA key for all tests, because generating it is the slow part.
func signingKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	keyOnce.Do(func() {
		var err error
		if testKey, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
			panic(err)
		}
	})
	return testKey
}

func certPEM(t *testing.T, public any, private crypto.Signer) []byte {
	t.Helper()
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "sns.amazonaws.com"},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, public, private)
	if err != nil {
		t.Fatal(err)
	}
	return pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
}

// transport answers every request with the same body and counts the requests, so no test
// touches the network.
type transport struct {
	status   int
	body     []byte
	requests atomic.Int32
}

func (tr *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	tr.requests.Add(1)
	return &http.Response{
		StatusCode: tr.status,
		Status:     http.StatusText(tr.status),
		Body:       io.NopCloser(bytes.NewReader(tr.body)),
		Header:     http.Header{},
		Request:    req,
	}, nil
}

func newVerifier(t *testing.T, topics ...string) (*Verifier, *transport) {
	t.Helper()
	tr := &transport{status: http.StatusOK, body: certPEM(t, &signingKey(t).PublicKey, signingKey(t))}
	v := NewVerifier(topics...)
	v.Client.Transport = tr
	return v, tr
}

// signed returns m with its signature set, made the way SNS makes it.
func signed(t *testing.T, m Message, version string) *Message {
	t.Helper()
	m.SignatureVersion = version
	if m.SigningCertURL == "" {
		m.SigningCertURL = testCertURL
	}
	toSign, err := m.stringToSign()
	if err != nil {
		t.Fatal(err)
	}
	var digest []byte
	var hash crypto.Hash
	if version == "1" {
		sum := sha1.Sum([]byte(toSign))
		digest, hash = sum[:], crypto.SHA1
	} else {
		sum := sha256.Sum256([]byte(toSign))
		digest, hash = sum[:], crypto.SHA256
	}
	sig, err := rsa.SignPKCS1v15(rand.Reader, signingKey(t), hash, digest)
	if err != nil {
		t.Fatal(err)
	}
	m.Signature = base64.StdEncoding.EncodeToString(sig)
	return &m
}

func notification() Message {
	return Message{
		Type:      TypeNotification,
		MessageID: "22b80b92-fdea-4c2c-8f9d-bdfb0c7bf324",
		TopicARN:  testTopic,
		Subject:   "My First Message",
		Message:   `{"notificationType":"Delivery","mail":{"messageId":"abc"}}`,
		Timestamp: "2024-05-01T10:30:00.123Z",
	}
}

func subscriptionConfirmation() Message {
	return Message{
		Type:         TypeSubscriptionConfirmation,
		MessageID:    "165545c9-2a5c-472c-8df2-7ff2be2b3b1b",
		Token:        "2336412f37fb687f5d51e6e2425f004a",
		TopicARN:     testTopic,
		Message:      "You have chosen to subscribe to the topic " + testTopic + ".\nTo confirm the subscription, visit the SubscribeURL included in this message.",
		SubscribeURL: "https://sns.ap-northeast-2.amazonaws.com/?Action=ConfirmSubscription&TopicArn=" + testTopic + "&Token=2336412f37fb687f5d51e6e2425f004a",
		Timestamp:    "2024-05-01T10:30:00.123Z",
	}
}

// TestStringToSign pins the text that is signed, byte for byte, to the format in the AWS
// documentation. The other tests sign with the same function, so they could not notice a
// mistake in it.
func TestStringToSign(t *testing.T) {
	tests := []struct {
		name string
		msg  Message
		want string
	}{
		{
			name: "notification with a subject",
			msg: Message{
				Type: TypeNotification, MessageID: "id-1", TopicARN: "arn:t", Subject: "Hi",
				Message: "Hello", Timestamp: "2024-01-01T00:00:00.000Z",
			},
			want: "Message\nHello\nMessageId\nid-1\nSubject\nHi\nTimestamp\n2024-01-01T00:00:00.000Z\nTopicArn\narn:t\nType\nNotification\n",
		},
		{
			name: "notification without a subject",
			msg: Message{
				Type: TypeNotification, MessageID: "id-1", TopicARN: "arn:t",
				Message: "Hello", Timestamp: "2024-01-01T00:00:00.000Z",
			},
			want: "Message\nHello\nMessageId\nid-1\nTimestamp\n2024-01-01T00:00:00.000Z\nTopicArn\narn:t\nType\nNotification\n",
		},
		{
			name: "subscription confirmation",
			msg: Message{
				Type: TypeSubscriptionConfirmation, MessageID: "id-2", TopicARN: "arn:t", Token: "tok",
				Message: "Confirm", SubscribeURL: "https://sns.example/confirm", Timestamp: "2024-01-01T00:00:00.000Z",
			},
			want: "Message\nConfirm\nMessageId\nid-2\nSubscribeURL\nhttps://sns.example/confirm\nTimestamp\n2024-01-01T00:00:00.000Z\nToken\ntok\nTopicArn\narn:t\nType\nSubscriptionConfirmation\n",
		},
		{
			name: "unsubscribe confirmation signs the same fields",
			msg: Message{
				Type: TypeUnsubscribeConfirmation, MessageID: "id-3", TopicARN: "arn:t", Token: "tok",
				Message: "Bye", SubscribeURL: "https://sns.example/resubscribe", Timestamp: "2024-01-01T00:00:00.000Z",
			},
			want: "Message\nBye\nMessageId\nid-3\nSubscribeURL\nhttps://sns.example/resubscribe\nTimestamp\n2024-01-01T00:00:00.000Z\nToken\ntok\nTopicArn\narn:t\nType\nUnsubscribeConfirmation\n",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.msg.stringToSign()
			if err != nil {
				t.Fatal(err)
			}
			if got != tc.want {
				t.Errorf("stringToSign:\n got %q\nwant %q", got, tc.want)
			}
		})
	}

	if _, err := (&Message{Type: "Other"}).stringToSign(); err == nil {
		t.Error("an unknown message type has no string to sign")
	}
}

func TestVerifyAccepts(t *testing.T) {
	for _, version := range []string{"1", "2"} {
		for name, msg := range map[string]Message{
			"notification":              notification(),
			"subscription confirmation": subscriptionConfirmation(),
		} {
			t.Run(name+" v"+version, func(t *testing.T) {
				v, _ := newVerifier(t, testTopic)
				if err := v.Verify(context.Background(), signed(t, msg, version)); err != nil {
					t.Fatalf("Verify: %v", err)
				}
			})
		}
	}

	t.Run("notification without a subject", func(t *testing.T) {
		v, _ := newVerifier(t, testTopic)
		m := notification()
		m.Subject = ""
		if err := v.Verify(context.Background(), signed(t, m, "2")); err != nil {
			t.Fatalf("Verify: %v", err)
		}
	})

	t.Run("one of several topics", func(t *testing.T) {
		v, _ := newVerifier(t, "arn:aws:sns:us-east-1:1:other", testTopic)
		if err := v.Verify(context.Background(), signed(t, notification(), "2")); err != nil {
			t.Fatalf("Verify: %v", err)
		}
	})
}

func TestVerifyRejects(t *testing.T) {
	ctx := context.Background()

	tamper := func(change func(*Message)) *Message {
		m := signed(t, notification(), "2")
		change(m)
		return m
	}

	tests := []struct {
		name string
		msg  *Message
	}{
		{"changed message", tamper(func(m *Message) { m.Message = `{"notificationType":"Bounce"}` })},
		{"changed subject", tamper(func(m *Message) { m.Subject = "other" })},
		{"changed timestamp", tamper(func(m *Message) { m.Timestamp = "2024-05-01T10:31:00.000Z" })},
		{"changed message id", tamper(func(m *Message) { m.MessageID = "another" })},
		{"changed type", tamper(func(m *Message) { m.Type = TypeSubscriptionConfirmation })},
		{"signature of another version", tamper(func(m *Message) { m.SignatureVersion = "1" })},
		{"unknown signature version", tamper(func(m *Message) { m.SignatureVersion = "3" })},
		{"no signature version", tamper(func(m *Message) { m.SignatureVersion = "" })},
		{"no signature", tamper(func(m *Message) { m.Signature = "" })},
		{"signature that is not base64", tamper(func(m *Message) { m.Signature = "%%%" })},
		{"unknown message type", tamper(func(m *Message) { m.Type = "Other" })},
		{"topic that is not configured", signed(t, func() Message {
			m := notification()
			m.TopicARN = "arn:aws:sns:ap-northeast-2:999999999999:someone-elses"
			return m
		}(), "2")},
		{"no topic", signed(t, func() Message { m := notification(); m.TopicARN = ""; return m }(), "2")},
		{"certificate on another host", tamper(func(m *Message) { m.SigningCertURL = "https://example.com/cert.pem" })},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			v, _ := newVerifier(t, testTopic)
			if err := v.Verify(ctx, tc.msg); err == nil {
				t.Fatal("Verify accepted the message")
			}
		})
	}

	t.Run("signed by another key", func(t *testing.T) {
		v, tr := newVerifier(t, testTopic)
		other, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			t.Fatal(err)
		}
		tr.body = certPEM(t, &other.PublicKey, other) // the URL serves a different certificate
		if err := v.Verify(ctx, signed(t, notification(), "2")); err == nil {
			t.Fatal("Verify accepted a signature that the certificate does not match")
		}
	})

	t.Run("no topic configured", func(t *testing.T) {
		v, _ := newVerifier(t)
		if v.Configured() {
			t.Error("a Verifier without topics reports that it is configured")
		}
		if err := v.Verify(ctx, signed(t, notification(), "2")); err == nil {
			t.Fatal("a Verifier without topics has to fail closed")
		}
	})

	t.Run("empty topic names do not configure it", func(t *testing.T) {
		if NewVerifier("", "").Configured() {
			t.Error("empty topic names count as configured")
		}
	})
}

func TestCertificateProblems(t *testing.T) {
	ctx := context.Background()
	m := signed(t, notification(), "2")

	t.Run("download fails", func(t *testing.T) {
		v, tr := newVerifier(t, testTopic)
		tr.status = http.StatusNotFound
		if err := v.Verify(ctx, m); err == nil {
			t.Fatal("Verify accepted a message although the certificate could not be downloaded")
		}
	})

	t.Run("not PEM", func(t *testing.T) {
		v, tr := newVerifier(t, testTopic)
		tr.body = []byte("not a certificate")
		if err := v.Verify(ctx, m); err == nil {
			t.Fatal("Verify accepted a certificate that is not PEM")
		}
	})

	t.Run("not an RSA key", func(t *testing.T) {
		v, tr := newVerifier(t, testTopic)
		ec, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Fatal(err)
		}
		tr.body = certPEM(t, &ec.PublicKey, ec)
		if err := v.Verify(ctx, m); err == nil {
			t.Fatal("Verify accepted a certificate without an RSA key")
		}
	})

	t.Run("a failed download is not cached", func(t *testing.T) {
		v, tr := newVerifier(t, testTopic)
		tr.status = http.StatusInternalServerError
		_ = v.Verify(ctx, m)
		tr.status = http.StatusOK
		if err := v.Verify(ctx, m); err != nil {
			t.Fatalf("Verify after the certificate became available: %v", err)
		}
	})

	t.Run("the certificate is downloaded once", func(t *testing.T) {
		v, tr := newVerifier(t, testTopic)
		for range 3 {
			if err := v.Verify(ctx, m); err != nil {
				t.Fatal(err)
			}
		}
		if got := tr.requests.Load(); got != 1 {
			t.Errorf("certificate requests = %d, want 1", got)
		}
	})

	t.Run("redirects are not followed", func(t *testing.T) {
		v := NewVerifier(testTopic)
		err := v.Client.CheckRedirect(nil, nil)
		if err != http.ErrUseLastResponse {
			t.Errorf("CheckRedirect = %v, want http.ErrUseLastResponse", err)
		}
	})
}

func TestCheckCertURL(t *testing.T) {
	good := []string{
		testCertURL,
		"https://sns.us-east-1.amazonaws.com/SimpleNotificationService-abc.pem",
		"https://sns.us-gov-west-1.amazonaws.com/SimpleNotificationService-abc.pem",
		"https://sns.cn-north-1.amazonaws.com.cn/SimpleNotificationService-abc.pem",
		"https://sns.eu-west-1.amazonaws.com:443/SimpleNotificationService-abc.pem",
		"https://SNS.EU-WEST-1.AMAZONAWS.COM/SimpleNotificationService-abc.pem",
	}
	for _, raw := range good {
		if _, err := checkCertURL(raw); err != nil {
			t.Errorf("checkCertURL(%q) = %v, want it accepted", raw, err)
		}
	}

	bad := []string{
		"",
		"not a url",
		"http://sns.us-east-1.amazonaws.com/cert.pem",
		"ftp://sns.us-east-1.amazonaws.com/cert.pem",
		"https://example.com/cert.pem",
		"https://amazonaws.com/cert.pem",
		"https://s3.us-east-1.amazonaws.com/cert.pem",
		"https://sns.us-east-1.amazonaws.com.evil.example/cert.pem",
		"https://evil-sns.us-east-1.amazonaws.com/cert.pem",
		"https://sns.us-east-1.amazonaws.com@evil.example/cert.pem",
		"https://sns.us-east-1.amazonaws.com:evil@evil.example/cert.pem",
		"https://user@sns.us-east-1.amazonaws.com/cert.pem",
		"https://sns.us-east-1.amazonaws.com:8443/cert.pem",
		"https://sns.us-east-1.amazonaws.com./cert.pem",
		"https://127.0.0.1/cert.pem",
		"https://[::1]/cert.pem",
		"https://sns.us-east-1.amazonaws.com/cert.txt",
		"https://sns.us-east-1.amazonaws.com/cert.pem?redirect=https://evil.example",
		"https://sns.us-east-1.amazonaws.com/cert.pem#x",
		"https://sns.us-east-1.amazonaws.com/",
		"//sns.us-east-1.amazonaws.com/cert.pem",
		"https:///sns.us-east-1.amazonaws.com/cert.pem",
		"https://sns.us-east-1.amazonaws.com\\@evil.example/cert.pem",
		"https://sns.us-east-1.amazonaws.com%2f@evil.example/cert.pem",
		strings.Repeat("a", 5000),
	}
	for _, raw := range bad {
		if _, err := checkCertURL(raw); err == nil {
			t.Errorf("checkCertURL(%q) accepted a URL that is not an SNS certificate", raw)
		}
	}
}

// A message without a signature, or from a topic that is not configured, is refused before any
// certificate is downloaded, so a caller cannot make the server send requests.
func TestRefusedBeforeDownloading(t *testing.T) {
	ctx := context.Background()

	unsigned := signed(t, notification(), "2")
	unsigned.Signature = ""

	foreign := signed(t, notification(), "2")
	foreign.TopicARN = "arn:aws:sns:ap-northeast-2:999999999999:attacker"

	for name, m := range map[string]*Message{"no signature": unsigned, "foreign topic": foreign} {
		v, tr := newVerifier(t, testTopic)
		if err := v.Verify(ctx, m); err == nil {
			t.Errorf("%s: Verify accepted the message", name)
		}
		if got := tr.requests.Load(); got != 0 {
			t.Errorf("%s: %d certificate requests, want none", name, got)
		}
	}
}
