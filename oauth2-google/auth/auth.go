package auth

import (
	"encoding/json"
	"fiber-oauth-google/config"
	"fiber-oauth-google/model"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func ConfigGoogle() *oauth2.Config {
	conf := &oauth2.Config{
		ClientID:     config.Config("Client"),
		ClientSecret: config.Config("Secret"),
		RedirectURL:  config.Config("redirect_url"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email"}, // you can use other scopes to get more data
		Endpoint: google.Endpoint,
	}
	return conf
}

func GetEmail(token string) string {
	reqURL, err := url.Parse("https://www.googleapis.com/oauth2/v1/userinfo")
	ptoken := fmt.Sprintf("Bearer %s", token)
	res := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Authorization": {ptoken}},
	}
	req, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)

	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var data model.GoogleResponse
	errorz := json.Unmarshal(body, &data)
	if errorz != nil {

		panic(errorz)
	}
	return data.Email
}
