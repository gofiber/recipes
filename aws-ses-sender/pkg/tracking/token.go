// Package tracking signs the open-tracking links that are put into the emails.
//
// The link of the tracking image holds the id of the email request, which is a counter that
// anyone can count along. Without a signature, anybody could report any email as opened. The
// token is an HMAC of the id, so only a link that the sender created is counted.
//
// The secret has to be a random string of its own. Every email carries a link, so everyone who
// gets one email can try guesses for the secret offline until the token matches. Do not use a
// password or an API key for it.
package tracking

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"html"
	"net/url"
	"strconv"
	"strings"
)

// ParseID reads the requestId of a tracking link, a plain decimal number. It parses to the size
// of uint (bit size 0), so converting the result to uint cannot truncate it on a 32-bit build,
// where a larger number is an error. An id that does not exist fails Valid later, because
// only the sender can sign it.
func ParseID(s string) (uint, error) {
	id, err := strconv.ParseUint(s, 10, 0)
	return uint(id), err
}

// Token returns the token for the link of the email request with the given id.
func Token(secret string, requestID uint) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte("open:" + strconv.FormatUint(uint64(requestID), 10)))
	return hex.EncodeToString(mac.Sum(nil))
}

// Valid reports whether token is the token of the request. An empty secret never validates,
// because anyone could compute the token for it.
func Valid(secret string, requestID uint, token string) bool {
	if secret == "" {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(Token(secret, requestID)), []byte(token)) == 1
}

// URL returns the link of the tracking image for the email request with the given id. The open
// event handler serves the link at /v1/events/open.
func URL(serverHost, secret string, requestID uint) string {
	query := url.Values{
		"requestId": {strconv.FormatUint(uint64(requestID), 10)},
		"token":     {Token(secret, requestID)},
	}
	return strings.TrimRight(serverHost, "/") + "/v1/events/open?" + query.Encode()
}

// Image returns the HTML of the tracking image for the email request, to be put at the end of
// the body. Without a secret the open event handler would accept no token, so there is no image:
// the result is empty.
func Image(serverHost, secret string, requestID uint) string {
	if secret == "" {
		return ""
	}
	return `<img src="` + html.EscapeString(URL(serverHost, secret, requestID)) + `">`
}
