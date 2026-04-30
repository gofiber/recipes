package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"fiber-oauth-google/config"
	"fiber-oauth-google/model"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOAuthConfig *oauth2.Config
	once              sync.Once
	httpClient        = &http.Client{Timeout: 10 * time.Second}
)

// ConfigGoogle returns the singleton OAuth2 config for Google.
func ConfigGoogle() *oauth2.Config {
	once.Do(func() {
		googleOAuthConfig = &oauth2.Config{
			ClientID:     config.Config("GOOGLE_CLIENT_ID"),
			ClientSecret: config.Config("GOOGLE_CLIENT_SECRET"),
			RedirectURL:  config.Config("GOOGLE_REDIRECT_URL"),
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
			}, // you can use other scopes to get more data
			Endpoint: google.Endpoint,
		}
	})
	return googleOAuthConfig
}

// GetEmail retrieves the email address of the authenticated user.
func GetEmail(token string) (string, error) {
	reqURL, err := url.Parse("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		return "", fmt.Errorf("failed to parse userinfo URL: %w", err)
	}
	ptoken := fmt.Sprintf("Bearer %s", token)
	res := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Authorization": {ptoken},
		},
	}
	req, err := httpClient.Do(res)
	if err != nil {
		return "", fmt.Errorf("failed to fetch userinfo: %w", err)
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read userinfo response: %w", err)
	}
	var data model.GoogleResponse
	if err = json.Unmarshal(body, &data); err != nil {
		return "", fmt.Errorf("failed to parse userinfo response: %w", err)
	}
	return data.Email, nil
}
