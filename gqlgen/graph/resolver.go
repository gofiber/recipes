package graph

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"

	"github.com/gofiber/recipes/gqlgen/chat"
)

//go:generate go tool gqlgen generate

// MaxConnSubscriptions is the open subscription limit per websocket connection. It keeps one socket from taking every hub slot.
const MaxConnSubscriptions = 10

// ErrTooManyConnSubscriptions is returned by messageAdded while a connection has MaxConnSubscriptions open.
var ErrTooManyConnSubscriptions = fmt.Errorf("at most %d subscriptions per connection", MaxConnSubscriptions)

// Resolver is the root resolver. The hub owns all state and locking.
type Resolver struct {
	Hub *chat.Hub
}

type connSubscriptionsKey struct{}

// WithConnSubscriptions returns a websocket connection context that counts the subscriptions opened on it.
func WithConnSubscriptions(ctx context.Context) context.Context {
	return context.WithValue(ctx, connSubscriptionsKey{}, new(atomic.Int32))
}

// acquireSubscription counts a subscription against its connection until ctx, the subscription context, is done.
func acquireSubscription(ctx context.Context) error {
	n, ok := ctx.Value(connSubscriptionsKey{}).(*atomic.Int32)
	if !ok {
		return errors.New("subscriptions need a websocket connection")
	}
	if n.Add(1) > MaxConnSubscriptions {
		n.Add(-1)
		return ErrTooManyConnSubscriptions
	}
	context.AfterFunc(ctx, func() { n.Add(-1) })
	return nil
}
