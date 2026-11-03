package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	fastws "github.com/fasthttp/websocket"
	"github.com/gofiber/contrib/v3/websocket"

	"github.com/gofiber/recipes/gqlgen/graph"
)

const (
	writeTimeout = 10 * time.Second
	closeTimeout = time.Second
)

// fiberWebsocket hands a connection upgraded by the Fiber websocket middleware to gqlgen's transport.Websocket.
type fiberWebsocket struct{}

// Accept returns the already upgraded connection carried by the response writer.
func (fiberWebsocket) Accept(w http.ResponseWriter, _ *http.Request, _ transport.WebsocketAcceptOptions) (transport.WebsocketConn, error) {
	uw, ok := w.(*upgradedWriter)
	if !ok {
		return nil, errors.New("gqlgen websocket: request was not upgraded by the Fiber websocket route")
	}
	return &wsConn{c: uw.conn}, nil
}

// upgradedWriter is the http.ResponseWriter passed to gqlgen. Writes are discarded because the socket is already upgraded.
type upgradedWriter struct {
	conn   *websocket.Conn
	header http.Header
}

func (w *upgradedWriter) Header() http.Header {
	if w.header == nil {
		w.header = http.Header{}
	}
	return w.header
}

func (*upgradedWriter) Write(p []byte) (int, error) { return len(p), nil }

func (*upgradedWriter) WriteHeader(int) {}

// wsConn wraps a Fiber websocket connection as a gqlgen transport.WebsocketConn.
type wsConn struct {
	c *websocket.Conn
}

var (
	_ transport.WebsocketConn          = (*wsConn)(nil)
	_ transport.WebsocketReadLimiter   = (*wsConn)(nil)
	_ transport.WebsocketReadDeadliner = (*wsConn)(nil)
)

func (w *wsConn) NextReader() (int, io.Reader, error) {
	t, r, err := w.c.NextReader()
	if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
		return t, nil, transport.ErrWebsocketClosed
	}
	return t, r, err
}

// WriteJSON closes the socket when a write fails. That ends gqlgen's read loop, which cancels the subscriptions.
func (w *wsConn) WriteJSON(v any) error {
	if err := w.c.SetWriteDeadline(time.Now().Add(writeTimeout)); err != nil {
		_ = w.c.Close()
		return err
	}
	if err := w.c.WriteJSON(v); err != nil {
		_ = w.c.Close()
		return err
	}
	return nil
}

func (w *wsConn) WriteClose(code int, reason string) error {
	return w.c.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(code, reason), time.Now().Add(closeTimeout))
}

// isDisconnect reports errors from a connection that is closing, gone or not reading.
// ErrCloseSent is the fasthttp/websocket sentinel. The contrib copy is a different value.
func isDisconnect(err error) bool {
	if errors.Is(err, net.ErrClosed) || errors.Is(err, fastws.ErrCloseSent) || errors.Is(err, syscall.ECONNRESET) || errors.Is(err, syscall.EPIPE) {
		return true
	}
	if _, ok := errors.AsType[*fastws.CloseError](err); ok {
		return true
	}
	ne, ok := errors.AsType[net.Error](err)
	return ok && ne.Timeout()
}

// logWebsocketError logs every websocket error except disconnects.
func logWebsocketError(_ context.Context, err error) {
	if isDisconnect(err) {
		return
	}
	log.Printf("graphql websocket: %v", err)
}

func (w *wsConn) Close() error { return w.c.Close() }

func (w *wsConn) Subprotocol() string { return w.c.Subprotocol() }

func (w *wsConn) SetReadLimit(limit int64) { w.c.SetReadLimit(limit) }

func (w *wsConn) SetReadDeadline(t time.Time) error { return w.c.SetReadDeadline(t) }

var (
	errShuttingDown = errors.New("server shutting down")
	errTooManyConns = errors.New("too many connections")
)

// connGroup counts live websocket handlers. It caps them at limit and lets shutdown wait for them.
type connGroup struct {
	limit   int
	mu      sync.Mutex
	wg      sync.WaitGroup
	open    int
	closing bool
}

// Add registers a connection. It fails once Wait has been called or while limit connections are open.
func (g *connGroup) Add() error {
	g.mu.Lock()
	defer g.mu.Unlock()
	if g.closing {
		return errShuttingDown
	}
	if g.open >= g.limit {
		return errTooManyConns
	}
	g.open++
	g.wg.Add(1)
	return nil
}

// Done releases a connection registered with Add.
func (g *connGroup) Done() {
	g.mu.Lock()
	g.open--
	g.mu.Unlock()
	g.wg.Done()
}

// Wait stops new registrations and waits for every registered connection or for ctx.
func (g *connGroup) Wait(ctx context.Context) error {
	g.mu.Lock()
	g.closing = true
	g.mu.Unlock()

	done := make(chan struct{})
	go func() {
		g.wg.Wait()
		close(done)
	}()
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// graphqlWebsocket runs gqlgen's websocket transport on a connection upgraded by Fiber.
func graphqlWebsocket(ctx context.Context, srv *handler.Server, conns *connGroup) func(*websocket.Conn) {
	return func(c *websocket.Conn) {
		if err := conns.Add(); err != nil {
			code := websocket.CloseGoingAway
			if errors.Is(err, errTooManyConns) {
				code = websocket.CloseTryAgainLater
			}
			closeConn(c, code, err.Error())
			return
		}
		defer conns.Done()
		// The middleware leaves the socket open on return. If gqlgen already sent a close frame,
		// such as 1002 on init timeout, the client sees only that one.
		defer closeConn(c, websocket.CloseNormalClosure, "terminated")

		// Canceling connCtx ends every subscription on this connection, because gqlgen keeps the request context.
		connCtx, cancel := context.WithCancel(ctx)
		defer cancel()
		connCtx = graph.WithConnSubscriptions(connCtx)

		// Cannot fail with a constant method and URL.
		req, _ := http.NewRequestWithContext(connCtx, http.MethodGet, "/query", http.NoBody)
		req.Header.Set("Upgrade", "websocket")
		srv.ServeHTTP(&upgradedWriter{conn: c}, req)
	}
}

// closeConn sends a close frame and closes the socket. A frame already sent by gqlgen wins.
func closeConn(c *websocket.Conn, code int, reason string) {
	_ = c.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(code, reason), time.Now().Add(closeTimeout))
	_ = c.Close()
}
