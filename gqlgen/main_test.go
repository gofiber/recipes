package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/gofiber/recipes/gqlgen/chat"
	"github.com/gofiber/recipes/gqlgen/graph"
)

type wsMessage struct {
	ID      string          `json:"id,omitempty"`
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload,omitempty"`
}

const (
	subscription = `subscription { messageAdded { id author text } }`
	postMutation = `mutation($in: NewMessage!) { postMessage(input: $in) { id author text } }`
)

// startServer runs the app on a free port with at most connLimit websocket connections.
// configure, if set, adjusts the gqlgen server before it serves.
func startServer(t *testing.T, connLimit int, configure func(*handler.Server)) (addr string, stop func() error, hub *chat.Hub) {
	t.Helper()
	ln, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	hub = chat.NewHub()
	srv := newGraphQLServer(hub)
	if configure != nil {
		configure(srv)
	}
	conns := &connGroup{limit: connLimit}
	app := newApp(ctx, srv, conns)

	errc := make(chan error, 1)
	go func() {
		errc <- app.Listener(ln, fiber.ListenConfig{DisableStartupMessage: true})
	}()

	stopped := false
	stop = func() error {
		if stopped {
			return nil
		}
		stopped = true
		cancel()
		err := shutdown(app, conns)
		select {
		case lerr := <-errc:
			if lerr != nil && !errors.Is(lerr, net.ErrClosed) {
				err = errors.Join(err, lerr)
			}
		case <-time.After(time.Second):
			err = errors.Join(err, errors.New("listener did not return"))
		}
		return err
	}
	t.Cleanup(func() { _ = stop() })
	return ln.Addr().String(), stop, hub
}

func graphqlPost(t *testing.T, addr, query string, vars map[string]any) string {
	t.Helper()
	body, err := json.Marshal(map[string]any{"query": query, "variables": vars})
	if err != nil {
		t.Fatal(err)
	}
	client := &http.Client{Transport: &http.Transport{DisableKeepAlives: true}, Timeout: 5 * time.Second}
	resp, err := client.Post("http://"+addr+"/query", "application/json", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = resp.Body.Close() }()
	out, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	return string(out)
}

// client reads frames in the background. A wait can then time out without breaking the socket.
type client struct {
	conn   *websocket.Conn
	msgs   chan wsMessage
	closed chan int // close code sent by the server, or -1 for any other read error
}

func dial(t *testing.T, addr string) *client {
	t.Helper()
	d := websocket.Dialer{Subprotocols: []string{"graphql-transport-ws"}, HandshakeTimeout: 5 * time.Second}
	c, resp, err := d.Dial("ws://"+addr+"/query", nil)
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()
	if c.Subprotocol() != "graphql-transport-ws" {
		t.Fatalf("negotiated %q", c.Subprotocol())
	}
	t.Cleanup(func() { _ = c.Close() })

	cl := &client{conn: c, msgs: make(chan wsMessage, 64), closed: make(chan int, 1)}
	go func() {
		for {
			var m wsMessage
			if err := c.ReadJSON(&m); err != nil {
				if ce, ok := errors.AsType[*websocket.CloseError](err); ok {
					cl.closed <- ce.Code
				} else {
					cl.closed <- -1
				}
				return
			}
			if m.Type != "ping" {
				cl.msgs <- m
			}
		}
	}()
	return cl
}

func (c *client) send(t *testing.T, m wsMessage) {
	t.Helper()
	if err := c.conn.WriteJSON(m); err != nil {
		t.Fatal(err)
	}
}

// next returns the next protocol message, or ok=false after wait.
func (c *client) next(wait time.Duration) (wsMessage, bool) {
	select {
	case m := <-c.msgs:
		return m, true
	case <-time.After(wait):
		return wsMessage{}, false
	}
}

func (c *client) expect(t *testing.T, typ, id, contains string) {
	t.Helper()
	m, ok := c.next(5 * time.Second)
	if !ok {
		t.Fatalf("no message, want type %q id %q", typ, id)
	}
	if m.Type != typ || m.ID != id || !strings.Contains(string(m.Payload), contains) {
		t.Fatalf("got %+v (payload %s), want type %q id %q containing %q", m, m.Payload, typ, id, contains)
	}
}

// expectClose waits for the server close frame and returns its code.
func (c *client) expectClose(t *testing.T) int {
	t.Helper()
	select {
	case code := <-c.closed:
		return code
	case <-time.After(2 * shutdownTimeout):
		t.Fatal("socket not closed")
		return 0
	}
}

func subscribeMsg(id string) wsMessage {
	return wsMessage{ID: id, Type: "subscribe", Payload: json.RawMessage(`{"query":"` + subscription + `"}`)}
}

func postHello(t *testing.T, addr string) string {
	t.Helper()
	var resp struct {
		Data struct {
			PostMessage struct{ ID string } `json:"postMessage"`
		} `json:"data"`
	}
	vars := map[string]any{"in": map[string]string{"author": "ana", "text": "hello"}}
	if err := json.Unmarshal([]byte(graphqlPost(t, addr, postMutation, vars)), &resp); err != nil {
		t.Fatal(err)
	}
	return resp.Data.PostMessage.ID
}

func helloEvent(id string) string {
	return `{"data":{"messageAdded":{"id":"` + id + `","author":"ana","text":"hello"}}}`
}

// waitFor polls cond until it holds or 5 seconds pass.
func waitFor(t *testing.T, what string, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for !cond() {
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for %s", what)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// lockedBuffer is a log destination that server goroutines and the test can share.
type lockedBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *lockedBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *lockedBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

// captureLog redirects the standard logger until the test ends.
func captureLog(t *testing.T) *lockedBuffer {
	t.Helper()
	b := &lockedBuffer{}
	prev := log.Writer()
	log.SetOutput(b)
	t.Cleanup(func() { log.SetOutput(prev) })
	return b
}

func TestGraphQLOverHTTPAndWebsocket(t *testing.T) {
	addr, stop, hub := startServer(t, maxConns, nil)

	// GraphiQL page.
	resp, err := http.Get("http://" + addr + "/")
	if err != nil {
		t.Fatal(err)
	}
	page, _ := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK || !strings.Contains(string(page), "graphiql") {
		t.Fatalf("GraphiQL: status %d", resp.StatusCode)
	}

	// A plain GET on /query skips the websocket middleware and reaches the adaptor route.
	resp, err = http.Get("http://" + addr + "/query?query=" + url.QueryEscape("{ messages { id } }"))
	if err != nil {
		t.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if got := string(body); got != `{"data":{"messages":[]}}` {
		t.Fatalf("GET query = %s", got)
	}

	if got := graphqlPost(t, addr, `{ messages(last: 0) { id } }`, nil); !strings.Contains(got, "last must be between 1 and 100") {
		t.Fatalf("messages(last: 0) = %s", got)
	}
	if got := graphqlPost(t, addr, postMutation, map[string]any{"in": map[string]string{"author": "ana", "text": ""}}); !strings.Contains(got, chat.ErrInvalidText.Error()) {
		t.Fatalf("empty text = %s", got)
	}

	ws := dial(t, addr)
	ws.send(t, wsMessage{Type: "connection_init"})
	ws.expect(t, "connection_ack", "", "")
	ws.send(t, subscribeMsg("1"))
	waitFor(t, "subscription 1", func() bool { return hub.Subscribers() == 1 })
	id := postHello(t, addr)
	ws.expect(t, "next", "1", helloEvent(id))

	// complete stops id 1: the next event reaches only the new id 2.
	ws.send(t, wsMessage{ID: "1", Type: "complete"})
	ws.expect(t, "complete", "1", "") // gqlgen confirms the stop
	waitFor(t, "subscription 1 to end", func() bool { return hub.Subscribers() == 0 })
	ws.send(t, subscribeMsg("2"))
	waitFor(t, "subscription 2", func() bool { return hub.Subscribers() == 1 })
	id = postHello(t, addr)
	ws.expect(t, "next", "2", helloEvent(id))
	if m, ok := ws.next(300 * time.Millisecond); ok {
		t.Fatalf("unexpected message after complete: %+v", m)
	}

	// A delivered message is already in the history.
	if got, want := graphqlPost(t, addr, `{ messages(last: 1) { id } }`, nil), `{"data":{"messages":[{"id":"`+id+`"}]}}`; got != want {
		t.Fatalf("messages = %s, want %s", got, want)
	}

	// Shutdown closes the socket with 1000 and waits for its handler.
	done := make(chan error, 1)
	start := time.Now()
	go func() { done <- stop() }()
	if code := ws.expectClose(t); code != websocket.CloseNormalClosure {
		t.Fatalf("close code %d, want 1000", code)
	}
	if err := <-done; err != nil {
		t.Fatalf("shutdown: %v", err)
	}
	if d := time.Since(start); d > shutdownTimeout {
		t.Fatalf("shutdown took %v", d)
	}
}

func TestConnGroupLimitAndWait(t *testing.T) {
	g := connGroup{limit: 1}
	if err := g.Add(); err != nil {
		t.Fatalf("first Add = %v", err)
	}
	if err := g.Add(); !errors.Is(err, errTooManyConns) {
		t.Fatalf("Add over the limit = %v", err)
	}
	g.Done()
	if err := g.Add(); err != nil {
		t.Fatalf("Add after Done = %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	if err := g.Wait(ctx); !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Wait with an open connection = %v", err)
	}
	if err := g.Add(); !errors.Is(err, errShuttingDown) {
		t.Fatalf("Add after Wait = %v", err)
	}
	g.Done()
	if err := g.Wait(context.Background()); err != nil {
		t.Fatalf("Wait after Done = %v", err)
	}
}

func TestQueryLimits(t *testing.T) {
	addr, _, _ := startServer(t, maxConns, nil)

	full := `messages(last: 100) { id author text createdAt }`
	if got := graphqlPost(t, addr, `{ `+full+` }`, nil); got != `{"data":{"messages":[]}}` {
		t.Fatalf("one full page = %s", got)
	}
	if got := graphqlPost(t, addr, `{ a: `+full+` b: `+full+` }`, nil); !strings.Contains(got, "operation has complexity 800, which exceeds the limit of 500") {
		t.Fatalf("two full pages = %s", got)
	}

	if got := graphqlPost(t, addr, `{ messages { id `+strings.Repeat("__typename ", maxTokens)+`} }`, nil); !strings.Contains(got, "exceeded token limit of 1000") {
		t.Fatalf("long query = %.200s", got)
	}

	// Each fragment spreads the previous one twice: 2^40 copies of F0 once expanded.
	var nested strings.Builder
	nested.WriteString(`{ messages(last: 100) { ...F40 } } fragment F0 on Message { id author text createdAt }`)
	for i := 1; i <= 40; i++ {
		fmt.Fprintf(&nested, " fragment F%d on Message { ...F%d ...F%d }", i, i-1, i-1)
	}
	start := time.Now()
	if got := graphqlPost(t, addr, nested.String(), nil); !strings.Contains(got, "which exceeds the limit of 500") {
		t.Fatalf("nested fragments = %.200s", got)
	}
	if d := time.Since(start); d > time.Second {
		t.Fatalf("nested fragments took %v", d)
	}

	// GraphiQL loads the schema with a full introspection query.
	if got := graphqlPost(t, addr, introspection.Query, nil); !strings.Contains(got, `"queryType":{"name":"Query"}`) {
		t.Fatalf("introspection = %.200s", got)
	}
	// Introspection is counted too: aliased nested type lists would otherwise build a huge response.
	aliases := func(n int, field string) string {
		var b strings.Builder
		for i := range n {
			fmt.Fprintf(&b, " a%d: %s", i, field)
		}
		return b.String()
	}
	attack := `{ __schema {` + aliases(20, "types { ...L1 }") + ` } }` +
		` fragment L1 on __Type {` + aliases(30, "fields { ...T }") + ` }` +
		` fragment T on __Field { type { ...L2 ofType { ...L2 ofType { ...L2 ofType { ...L2 } } } } }` +
		` fragment L2 on __Type {` + aliases(30, "fields { name }") + ` }`
	if got := graphqlPost(t, addr, attack, nil); !strings.Contains(got, "operation has complexity 735, which exceeds the limit of 500") {
		t.Fatalf("aliased introspection = %.200s", got)
	}

	// One operation may post 3 messages, fewer than a subscriber buffer holds.
	post := `postMessage(input: {author: "ana", text: "hi"}) { id }`
	if got := graphqlPost(t, addr, `mutation {`+aliases(3, post)+` }`, nil); strings.Contains(got, "errors") {
		t.Fatalf("3 posts = %s", got)
	}
	if got := graphqlPost(t, addr, `mutation {`+aliases(4, post)+` }`, nil); !strings.Contains(got, "operation has complexity 504, which exceeds the limit of 500") {
		t.Fatalf("4 posts = %s", got)
	}
}

func TestSubscriptionOverHTTPRefused(t *testing.T) {
	addr, _, hub := startServer(t, maxConns, nil)
	if got := graphqlPost(t, addr, subscription, nil); !strings.Contains(got, "subscriptions need a websocket connection") {
		t.Fatalf("POST subscription = %s", got)
	}
	if got := hub.Subscribers(); got != 0 {
		t.Fatalf("subscribers = %d, want 0", got)
	}
}

func TestClientDisconnectUnsubscribes(t *testing.T) {
	disconnects := map[string]func(*client) error{
		"close frame": func(c *client) error {
			return c.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(time.Second))
		},
		"close code 4499": func(c *client) error {
			return c.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(4499, ""), time.Now().Add(time.Second))
		},
		"dropped connection": func(c *client) error {
			return c.conn.UnderlyingConn().Close()
		},
		"connection reset": func(c *client) error {
			tcp := c.conn.UnderlyingConn().(*net.TCPConn)
			if err := tcp.SetLinger(0); err != nil {
				return err
			}
			return tcp.Close()
		},
	}
	for name, disconnect := range disconnects {
		t.Run(name, func(t *testing.T) {
			logs := captureLog(t)
			addr, _, hub := startServer(t, maxConns, nil)

			ws := dial(t, addr)
			ws.send(t, wsMessage{Type: "connection_init"})
			ws.expect(t, "connection_ack", "", "")
			ws.send(t, subscribeMsg("1"))
			waitFor(t, "the subscription to start", func() bool { return hub.Subscribers() == 1 })
			ws.expect(t, "next", "1", helloEvent(postHello(t, addr)))

			if err := disconnect(ws); err != nil {
				t.Fatal(err)
			}
			// gqlgen cancels the subscription context, then the hub removes the subscriber.
			waitFor(t, "the hub to drop the subscriber", func() bool { return hub.Subscribers() == 0 })
			postHello(t, addr)
			if got := logs.String(); got != "" {
				t.Fatalf("disconnect logged %q", got)
			}
		})
	}
}

func TestEvictionCompletesSubscription(t *testing.T) {
	// The gate holds the first subscription response, like a client too slow to read.
	gate := make(chan struct{})
	slowSubscriptions := func(srv *handler.Server) {
		srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
			if graphql.GetOperationContext(ctx).Operation.Operation == ast.Subscription {
				select {
				case <-gate:
				case <-ctx.Done():
				}
			}
			return next(ctx)
		})
	}
	addr, _, hub := startServer(t, maxConns, slowSubscriptions)

	ws := dial(t, addr)
	ws.send(t, wsMessage{Type: "connection_init"})
	ws.expect(t, "connection_ack", "", "")
	ws.send(t, subscribeMsg("1"))
	waitFor(t, "the subscription to start", func() bool { return hub.Subscribers() == 1 })

	// The buffer takes SubscriberBuffer posts; the next one evicts the subscriber.
	for range chat.SubscriberBuffer + 1 {
		postHello(t, addr)
	}
	if got := hub.Subscribers(); got != 0 {
		t.Fatalf("subscribers = %d, want 0 after eviction", got)
	}

	// The buffered messages are delivered, then the server ends the subscription with complete.
	close(gate)
	for i := 1; i <= chat.SubscriberBuffer; i++ {
		ws.expect(t, "next", "1", helloEvent(fmt.Sprint(i)))
	}
	ws.expect(t, "complete", "1", "")
}

func TestShutdownClosesSocketWithoutInit(t *testing.T) {
	addr, stop, _ := startServer(t, maxConns, nil)
	d := websocket.Dialer{Subprotocols: []string{"graphql-transport-ws"}}
	c, _, err := d.Dial("ws://"+addr+"/query", nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = c.Close() })

	// Shut down on the pong. gqlgen answers pings only after it registers the connection.
	start := time.Now()
	done := make(chan error, 1)
	c.SetPongHandler(func(string) error {
		go func() { done <- stop() }()
		return nil
	})
	if err := c.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Second)); err != nil {
		t.Fatal(err)
	}
	_ = c.SetReadDeadline(time.Now().Add(2 * shutdownTimeout))
	// The init timeout sends 1002 before shutdown returns.
	if _, _, err := c.ReadMessage(); !websocket.IsCloseError(err, websocket.CloseProtocolError) {
		t.Fatalf("read = %v, want close code %d", err, websocket.CloseProtocolError)
	}
	select {
	case err := <-done:
		if d := time.Since(start); err != nil || d > shutdownTimeout {
			t.Fatalf("shutdown = %v after %v", err, d)
		}
	case <-time.After(shutdownTimeout):
		t.Fatal("shutdown did not return")
	}
}

func TestHTTPRequestWithUpgradeHeader(t *testing.T) {
	logs := captureLog(t)
	addr, _, _ := startServer(t, maxConns, nil)
	req, err := http.NewRequest(http.MethodPost, "http://"+addr+"/query", strings.NewReader(`{"query":"{ messages { id } }"}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Upgrade", "h2c") // sent by curl --http2
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if resp.StatusCode != http.StatusOK || string(body) != `{"data":{"messages":[]}}` {
		t.Fatalf("status %d, body %s", resp.StatusCode, body)
	}
	if got := logs.String(); got != "" {
		t.Fatalf("logged %q", got)
	}
}

func TestOtherSubprotocolsRefused(t *testing.T) {
	logs := captureLog(t)
	addr, _, _ := startServer(t, maxConns, nil)
	for _, protocols := range [][]string{{"graphql-ws"}, nil} {
		d := websocket.Dialer{Subprotocols: protocols, HandshakeTimeout: 5 * time.Second}
		c, resp, err := d.Dial("ws://"+addr+"/query", nil)
		if err == nil {
			_ = c.Close()
			t.Fatalf("%q: upgraded, want refusal", protocols)
		}
		if resp == nil || resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("%q: response %v, err %v, want 400", protocols, resp, err)
		}
		_ = resp.Body.Close()
	}
	if got := logs.String(); got != "" {
		t.Fatalf("logged %q", got)
	}
}

func TestReusedSubscriptionIDsEndOnDisconnect(t *testing.T) {
	addr, _, hub := startServer(t, maxConns, nil)
	ws := dial(t, addr)
	ws.send(t, wsMessage{Type: "connection_init"})
	ws.expect(t, "connection_ack", "", "")
	// gqlgen keeps only the last subscription per id. The others end only with the connection.
	for range graph.MaxConnSubscriptions {
		ws.send(t, subscribeMsg("1"))
	}
	waitFor(t, "every subscription", func() bool { return hub.Subscribers() == graph.MaxConnSubscriptions })
	_ = ws.conn.Close()
	waitFor(t, "the hub to drop every subscriber", func() bool { return hub.Subscribers() == 0 })
}

func TestSubscriptionsPerConnectionLimit(t *testing.T) {
	addr, _, hub := startServer(t, maxConns, nil)
	ws := dial(t, addr)
	ws.send(t, wsMessage{Type: "connection_init"})
	ws.expect(t, "connection_ack", "", "")
	for i := range graph.MaxConnSubscriptions {
		ws.send(t, subscribeMsg(fmt.Sprint(i)))
	}
	waitFor(t, "every subscription", func() bool { return hub.Subscribers() == graph.MaxConnSubscriptions })

	ws.send(t, subscribeMsg("over"))
	ws.expect(t, "next", "over", graph.ErrTooManyConnSubscriptions.Error())
	ws.expect(t, "complete", "over", "")

	// Completing one subscription frees its slot.
	ws.send(t, wsMessage{ID: "0", Type: "complete"})
	ws.expect(t, "complete", "0", "")
	ws.send(t, subscribeMsg("again"))
	waitFor(t, "the new subscription", func() bool { return hub.Subscribers() == graph.MaxConnSubscriptions })
	if m, ok := ws.next(300 * time.Millisecond); ok {
		t.Fatalf("unexpected message: %+v (payload %s)", m, m.Payload)
	}
}

func TestConnectionLimit(t *testing.T) {
	addr, _, _ := startServer(t, 1, nil)
	first := dial(t, addr)
	first.send(t, wsMessage{Type: "connection_init"})
	first.expect(t, "connection_ack", "", "")
	if code := dial(t, addr).expectClose(t); code != websocket.CloseTryAgainLater {
		t.Fatalf("close code %d, want %d", code, websocket.CloseTryAgainLater)
	}
	first.send(t, wsMessage{ID: "1", Type: "subscribe", Payload: json.RawMessage(`{"query":"{ messages { id } }"}`)})
	first.expect(t, "next", "1", `{"data":{"messages":[]}}`)
}

func TestInvalidMessageClosesWithFrame(t *testing.T) {
	logs := captureLog(t)
	addr, _, _ := startServer(t, maxConns, nil)
	ws := dial(t, addr)
	ws.send(t, wsMessage{Type: "connection_init"})
	ws.expect(t, "connection_ack", "", "")
	ws.send(t, wsMessage{Type: "bogus"})
	if code := ws.expectClose(t); code != websocket.CloseNormalClosure {
		t.Fatalf("close code %d, want %d", code, websocket.CloseNormalClosure)
	}
	if got := strings.Count(logs.String(), "\n"); got > 1 {
		t.Fatalf("logged %d lines: %q", got, logs.String())
	}
}

func TestWebsocketErrorLogging(t *testing.T) {
	logs := captureLog(t)
	ctx := context.Background()
	for _, err := range []error{
		&websocket.CloseError{Code: 4499},
		&net.OpError{Op: "read", Err: os.NewSyscallError("read", syscall.ECONNRESET)},
		&net.OpError{Op: "write", Err: os.ErrDeadlineExceeded},
	} {
		logWebsocketError(ctx, transport.WebsocketError{Err: err})
	}
	if got := logs.String(); got != "" {
		t.Fatalf("disconnects logged %q", got)
	}
	for range 3 {
		logWebsocketError(ctx, errors.New("other"))
	}
	if got := strings.Count(logs.String(), "\n"); got != 3 {
		t.Fatalf("logged %d lines, want 3: %q", got, logs.String())
	}
}
