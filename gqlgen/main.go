package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/gofiber/recipes/gqlgen/chat"
	"github.com/gofiber/recipes/gqlgen/graph"
	"github.com/gofiber/recipes/gqlgen/graph/model"
)

const (
	// shutdownTimeout covers both the HTTP drain and the websocket wait.
	shutdownTimeout = 5 * time.Second
	// initTimeout must stay below shutdownTimeout. A socket that never sends connection_init then closes before the shutdown deadline.
	initTimeout = 3 * time.Second

	// maxComplexity fits one messages(last: 100) with all four Message fields.
	maxComplexity = 500
	// maxTokens is about five times the size of a full introspection query.
	maxTokens = 1000

	maxConns = 10_000

	wsSubprotocol = "graphql-transport-ws"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	conns := &connGroup{limit: maxConns}
	app := newApp(ctx, newGraphQLServer(chat.NewHub()), conns)

	errc := make(chan error, 1)
	go func() { errc <- app.Listen(":3000") }()
	select {
	case err := <-errc:
		log.Fatal(err)
	case <-ctx.Done():
		// A second Ctrl+C now kills the process at once.
		stop()
	}

	// gqlgen is already closing the initialized websocket connections, since ctx is done.
	if err := shutdown(app, conns); err != nil {
		log.Printf("shutdown: %v", err)
	}
}

// shutdown drains HTTP requests, then waits for websocket handlers, within one shutdownTimeout.
func shutdown(app *fiber.App, conns *connGroup) error {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		return fmt.Errorf("http: %w", err)
	}
	if err := conns.Wait(ctx); err != nil {
		return fmt.Errorf("websocket connections still open: %w", err)
	}
	return nil
}

// newGraphQLServer builds the gqlgen server shared by the HTTP and websocket routes.
func newGraphQLServer(hub *chat.Hub) *handler.Server {
	cfg := graph.Config{Resolvers: &graph.Resolver{Hub: hub}}
	// The default cost ignores list sizes. Charge the selection once per returned message.
	cfg.Complexity.Query.Messages = func(childComplexity int, last int32) int {
		return min(max(int(last), 1), chat.HistorySize) * childComplexity
	}
	// At more than maxComplexity/4 per post, one operation holds at most 3 posts, fewer than a subscriber buffer.
	cfg.Complexity.Mutation.PostMessage = func(c int, _ model.NewMessage) int { return maxComplexity/4 + c }
	srv := handler.New(graph.NewExecutableSchema(cfg))

	// Only bridged websocket requests arrive with an Upgrade header. The HTTP route strips it.
	srv.AddTransport(transport.Websocket{
		Implementation:   fiberWebsocket{},
		InitTimeout:      initTimeout,
		PingPongInterval: 10 * time.Second,
		MissingPongOk:    true, // keep manual clients such as websocat connected
		ErrorFunc:        logWebsocketError,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.SetParserTokenLimit(maxTokens)
	srv.Use(&complexityLimit{limit: maxComplexity})
	srv.Use(extension.Introspection{})
	return srv
}

// newApp registers GraphiQL and the /query routes. Canceling ctx closes every initialized websocket connection.
func newApp(ctx context.Context, srv *handler.Server, conns *connGroup) *fiber.App {
	app := fiber.New()

	app.Get("/", adaptor.HTTPHandlerFunc(playground.Handler("GraphQL", "/query")))

	// Next skips plain requests, which fall through to the HTTP route below.
	app.Get("/query", requireSubprotocol, websocket.New(graphqlWebsocket(ctx, srv, conns), websocket.Config{
		Next: func(c fiber.Ctx) bool {
			return !websocket.IsWebSocketUpgrade(c)
		},
		Subprotocols: []string{wsSubprotocol},
	}))
	// gqlgen treats any request with an Upgrade header, h2c included, as a websocket request.
	httpHandler := adaptor.HTTPHandler(srv)
	app.All("/query", func(c fiber.Ctx) error {
		c.Request().Header.Del(fiber.HeaderUpgrade)
		return httpHandler(c)
	})

	return app
}

// requireSubprotocol refuses upgrades that do not offer graphql-transport-ws, which gqlgen would serve as legacy graphql-ws.
// It reads the first Sec-WebSocket-Protocol header, as the upgrader does.
func requireSubprotocol(c fiber.Ctx) error {
	if !websocket.IsWebSocketUpgrade(c) {
		return c.Next()
	}
	for p := range strings.SplitSeq(c.Get(fiber.HeaderSecWebSocketProtocol), ",") {
		if strings.TrimSpace(p) == wsSubprotocol {
			return c.Next()
		}
	}
	return fiber.NewError(fiber.StatusBadRequest, "websocket subprotocol "+wsSubprotocol+" required")
}
