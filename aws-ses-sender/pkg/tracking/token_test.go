package tracking

import (
	"net/url"
	"strconv"
	"strings"
	"testing"
)

// The tokens below were made with Python (hmac.new(secret, b"open:<id>", sha256).hexdigest()), so
// they pin the format of what is signed and not only that two calls agree.
func TestTokenVectors(t *testing.T) {
	tests := []struct {
		secret string
		id     uint
		want   string
	}{
		{"test-secret", 42, "acb501cab463a4ec45b9a21124e72a130626484032a067d845cf67cb67cd525b"},
		{"test-secret", 7, "2353e12fccbf1c662591abbf7adda9016201eac5c695cd98c225b5238a50b184"},
		{"another-secret", 42, "13e6690699251ff4025bfe854fd53bf1df60959e0cae49db71dfee325d180f39"},
	}
	for _, tc := range tests {
		if got := Token(tc.secret, tc.id); got != tc.want {
			t.Errorf("Token(%q, %d) = %s, want %s", tc.secret, tc.id, got, tc.want)
		}
	}
}

func TestToken(t *testing.T) {
	const secret = "api-key"

	token := Token(secret, 42)
	if len(token) != 64 || strings.Trim(token, "0123456789abcdef") != "" {
		t.Fatalf("Token = %q, want 64 lower-case hex digits", token)
	}
	if token != Token(secret, 42) {
		t.Error("Token is not deterministic")
	}
	if token == Token(secret, 43) {
		t.Error("two requests share a token")
	}
	if token == Token("other-key", 42) {
		t.Error("two secrets share a token")
	}
}

func TestValid(t *testing.T) {
	const secret = "api-key"
	good := Token(secret, 42)

	tests := []struct {
		name   string
		secret string
		id     uint
		token  string
		want   bool
	}{
		{"the token of the request", secret, 42, good, true},
		{"the token of another request", secret, 43, good, false},
		{"a token made with another secret", secret, 42, Token("other-key", 42), false},
		{"an empty token", secret, 42, "", false},
		{"a truncated token", secret, 42, good[:32], false},
		{"a token with extra characters", secret, 42, good + "0", false},
		{"an upper-case token", secret, 42, strings.ToUpper(good), false},
		{"no secret configured", "", 42, Token("", 42), false},
		{"no secret and no token", "", 42, "", false},
	}
	for _, tc := range tests {
		if got := Valid(tc.secret, tc.id, tc.token); got != tc.want {
			t.Errorf("%s: Valid = %v, want %v", tc.name, got, tc.want)
		}
	}
}

func TestURL(t *testing.T) {
	got := URL("https://mail.example.com/", "api-key", 42)
	want := "https://mail.example.com/v1/events/open?requestId=42&token=" + Token("api-key", 42)
	if got != want {
		t.Errorf("URL = %q, want %q", got, want)
	}

	// The link has to give back what the open event handler needs to validate it.
	u, err := url.Parse(got)
	if err != nil {
		t.Fatal(err)
	}
	id, err := strconv.ParseUint(u.Query().Get("requestId"), 10, 63)
	if err != nil {
		t.Fatal(err)
	}
	if !Valid("api-key", uint(id), u.Query().Get("token")) {
		t.Error("the token in the link does not validate for the id in the link")
	}
}

func TestImage(t *testing.T) {
	got := Image("https://mail.example.com", "test-secret", 42)
	want := `<img src="https://mail.example.com/v1/events/open?requestId=42&amp;token=acb501cab463a4ec45b9a21124e72a130626484032a067d845cf67cb67cd525b">`
	if got != want {
		t.Errorf("Image = %q, want %q", got, want)
	}

	// There is no token without a secret, so the email gets no image.
	if got := Image("https://mail.example.com", "", 42); got != "" {
		t.Errorf("Image without a secret = %q, want nothing", got)
	}
}
