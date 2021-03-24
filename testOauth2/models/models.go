package models

import (
	"github.com/antigloss/go/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// {"error":"bad_verification_code","error_description":"The code passed is incorrect or expired.","error_uri":"https://docs.github.com/apps/managing-oauth-apps/troubleshooting-oauth-app-access-token-request-errors/#bad-verification-code"}
// {"access_token":"some value","token_type":"bearer","scope":""}

type OAuthAccessResponse struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	Scope            string `json:"scope"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorURI         string `json:"error_uri"`
}

// SYSLOG ...
var SYSLOG *logger.Logger

// ClientID ...
var ClientID string

// ClientSecret ...
var ClientSecret string

var MySessionStore *session.Store
