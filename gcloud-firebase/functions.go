package app

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp/fasthttputil"
)

// CloudFunctionRouteToFiber route cloud function http.Handler to *fiber.App
// Internally, google calls the function with the /execute base URL
func CloudFunctionRouteToFiber(fiberApp *fiber.App, w http.ResponseWriter, r *http.Request) error {
	return RouteToFiber(fiberApp, w, r, "/execute")
}

// RouteToFiber route http.Handler to *fiber.App
func RouteToFiber(fiberApp *fiber.App, w http.ResponseWriter, r *http.Request, rootURL ...string) error {
	ln := fasthttputil.NewInmemoryListener()
	defer ln.Close()

	// Copy request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s://%s%s", "http", "0.0.0.0", r.RequestURI)
	if len(rootURL) > 0 {
		url = strings.Replace(url, rootURL[0], "", -1)
	}

	proxyReq, err := http.NewRequest(r.Method, url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	proxyReq.Header = r.Header

	// Create http client
	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return ln.Dial()
			},
		},
	}

	// Serve request to internal HTTP client
	go func() {
		err := fiberApp.Listener(ln)
		if err != nil {
			log.Fatalf("server err : %v", err)
			panic(err)
		}
	}()

	// Call internal Fiber API
	response, err := client.Do(proxyReq)
	if err != nil {
		return err
	}

	// Copy response and headers
	for k, values := range response.Header {
		for _, v := range values {
			w.Header().Set(k, v)
		}
	}
	w.WriteHeader(response.StatusCode)

	io.Copy(w, response.Body)
	response.Body.Close()

	return nil
}
