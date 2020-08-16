package services

import (
	"fmt"
	. "github.com/itsursujit/fiber-boilerplate/app"
	"github.com/itsursujit/fiber-boilerplate/http"
	"io/ioutil"
)

type Webhook struct {
	Url          string
	Payload      interface{}
	RetryMax     int
	RetryWaitMax int
	Timeout      int
}

func (w *Webhook) Dispatch() {
	retryClient := http.NewClient(http.Options{
		RetryWaitMin:  1,
		RetryWaitMax:  1,
		Timeout:       10,
		RetryMax:      w.RetryMax,
		RespReadLimit: 0,
		KillIdleConn:  true,
		MaxPoolSize:   100,
		ReqPerSec:     100,
		Verbose:       true,
	})
	Log.Info().Msg("Webhook client initiated")
	go Worker(retryClient, w.Url, w.Payload)
}

func Worker(retryClient *http.Client, url string, payload interface{}) {
	Log.Info().Msg("Sending callback to webhook:" + url)
	resp, err := retryClient.Get(url)
	if err != nil {
		Log.Fatal().Err(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	Log.Info().Msg(fmt.Sprintf("Response Received: %s", bodyBytes))
}
