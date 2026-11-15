// Package tracking signs the open-tracking links that are put into the emails.
//
// The link of the tracking image holds the id of the email request, which is a counter that
// anyone can count along. Without a signature, anybody could report any email as opened. The
// token is an HMAC of the id, so only a link that the sender created is counted.
package tracking

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
)

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
