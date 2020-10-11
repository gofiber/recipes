// Package adapter ...
// author : @fenny (author of fiber)
// Package fiberadapter adds Fiber support for the aws-severless-go-api library.
// Uses the core package behind the scenes and exposes the New method to
// get a new instance and Proxy method to send request to the Fiber app.
package adapter

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/awslabs/aws-lambda-go-api-proxy/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"github.com/valyala/fasthttp"
)

// FiberLambda makes it easy to send API Gateway proxy events to a fiber.App.
// The library transforms the proxy event into an HTTP request and then
// creates a proxy response object from the *fiber.Ctx
type FiberLambda struct {
	core.RequestAccessor
	app *fiber.App
}

// New creates a new instance of the FiberLambda object.
// Receives an initialized *fiber.App object - normally created with fiber.New().
// It returns the initialized instance of the FiberLambda object.
func New(app *fiber.App) *FiberLambda {
	return &FiberLambda{
		app: app,
	}
}

// Proxy receives an API Gateway proxy event, transforms it into an http.Request
// object, and sends it to the fiber.App for routing.
// It returns a proxy response object generated from the http.ResponseWriter.
func (f *FiberLambda) Proxy(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fiberRequest, err := f.ProxyEventToHTTPRequest(req)
	return f.proxyInternal(fiberRequest, err)
}

// ProxyWithContext receives context and an API Gateway proxy event,
// transforms them into an http.Request object, and sends it to the echo.Echo for routing.
// It returns a proxy response object generated from the http.ResponseWriter.
func (f *FiberLambda) ProxyWithContext(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fiberRequest, err := f.EventToRequestWithContext(ctx, req)
	return f.proxyInternal(fiberRequest, err)
}

func (f *FiberLambda) proxyInternal(req *http.Request, err error) (events.APIGatewayProxyResponse, error) {

	if err != nil {
		return core.GatewayTimeout(), core.NewLoggedError("Could not convert proxy event to request: %v", err)
	}

	resp := core.NewProxyResponseWriter()
	f.adaptor(resp, req)

	proxyResponse, err := resp.GetProxyResponse()
	if err != nil {
		return core.GatewayTimeout(), core.NewLoggedError("Error while generating proxy response: %v", err)
	}

	return proxyResponse, nil
}

func (f *FiberLambda) adaptor(w http.ResponseWriter, r *http.Request) {
	// New fasthttp request
	var req fasthttp.Request
	// Convert net/http -> fasthttp request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, utils.StatusMessage(fiber.StatusInternalServerError), fiber.StatusInternalServerError)
		return
	}
	req.Header.SetMethod(r.Method)
	req.SetRequestURI(r.RequestURI)
	req.Header.SetContentLength(len(body))
	req.SetHost(r.Host)
	for key, val := range r.Header {
		for _, v := range val {
			req.Header.Add(key, v)
		}
	}
	_, _ = req.BodyWriter().Write(body)
	remoteAddr, err := net.ResolveTCPAddr("tcp", r.RemoteAddr)
	if err != nil {
		http.Error(w, utils.StatusMessage(fiber.StatusInternalServerError), fiber.StatusInternalServerError)
		return
	}

	// New fasthttp Ctx
	var fctx fasthttp.RequestCtx
	fctx.Init(&req, remoteAddr, nil)

	// Pass RequestCtx to Fiber router
	f.app.Handler()(&fctx)
	// Convert fasthttp Ctx > net/http
	fctx.Response.Header.VisitAll(func(k, v []byte) {
		sk := string(k)
		sv := string(v)
		w.Header().Set(sk, sv)
	})
	w.WriteHeader(fctx.Response.StatusCode())
	_, _ = w.Write(fctx.Response.Body())
}
