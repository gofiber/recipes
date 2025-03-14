package app

import (
	"io"
	"net"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

// CloudFunctionRouteToFiber route cloud function http.Handler to *fiber.App
// Internally, google calls the function with the /execute base URL
func CloudFunctionRouteToFiber(fiberApp *fiber.App, w http.ResponseWriter, r *http.Request) error {

	// Convert net/http -> fasthttp Ctx
	ctx := ConvertNetHTTPRequestToFastHTTPCtx(r, w)

	// Run Fiber
	fiberApp.Handler()(ctx)

	// Convert fasthttp Ctx -> net/http
	ctx.Response.Header.VisitAll(func(k, v []byte) {
		w.Header().Add(string(k), string(v))
	})
	w.WriteHeader(ctx.Response.StatusCode())
	_, err := w.Write(ctx.Response.Body())

	return err
}

// ConvertNetHTTPRequestToFastHTTPCtx converts a net/http.Request to fasthttp.RequestCtx
func ConvertNetHTTPRequestToFastHTTPCtx(r *http.Request, w http.ResponseWriter) *fasthttp.RequestCtx {
	// New fasthttp request
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	// Convert net/http -> fasthttp request
	if r.Body != nil {
		n, err := io.Copy(req.BodyWriter(), r.Body)
		req.Header.SetContentLength(int(n))

		if err != nil {
			http.Error(w, utils.StatusMessage(fiber.StatusInternalServerError), fiber.StatusInternalServerError)
			return nil
		}
	}
	req.Header.SetMethod(r.Method)
	req.SetRequestURI(r.RequestURI)
	req.SetHost(r.Host)
	req.Header.SetHost(r.Host)
	for key, val := range r.Header {
		for _, v := range val {
			req.Header.Set(key, v)
		}
	}

	if _, _, err := net.SplitHostPort(r.RemoteAddr); err != nil && err.(*net.AddrError).Err == "missing port in address" {
		r.RemoteAddr = net.JoinHostPort(r.RemoteAddr, "80")
	}
	remoteAddr, err := net.ResolveTCPAddr("tcp", r.RemoteAddr)
	if err != nil {
		http.Error(w, utils.StatusMessage(fiber.StatusInternalServerError), fiber.StatusInternalServerError)
		return nil
	}

	var fctx fasthttp.RequestCtx
	fctx.Init(req, remoteAddr, nil)

	return &fctx
}
