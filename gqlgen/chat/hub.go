// Package chat keeps a bounded message history and fans new messages out to subscribers.
package chat

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/recipes/gqlgen/graph/model"
)

const (
	// HistorySize is the number of messages kept for the messages query.
	HistorySize = 100
	// SubscriberBuffer is how many messages a subscriber may have pending before it is evicted.
	SubscriberBuffer = 16
	// MaxAuthorLen is the author length limit in runes, after trimming.
	MaxAuthorLen = 32
	// MaxTextLen is the text length limit in runes, after trimming.
	MaxTextLen = 280
	// MaxSubscribers is the open subscription limit. Every Post sends to all of them under the lock.
	MaxSubscribers = 10_000
)

var (
	// ErrInvalidAuthor is returned for an empty or too long author.
	ErrInvalidAuthor = fmt.Errorf("author must be 1 to %d characters", MaxAuthorLen)
	// ErrInvalidText is returned for an empty or too long text.
	ErrInvalidText = fmt.Errorf("text must be 1 to %d characters", MaxTextLen)
	// ErrTooManySubscribers is returned by Subscribe once MaxSubscribers subscriptions are open.
	ErrTooManySubscribers = errors.New("too many subscribers")
)

// Hub stores recent messages and delivers new ones to subscribers in id order.
type Hub struct {
	mu      sync.RWMutex
	history [HistorySize]*model.Message // message id n is stored at (n-1) % HistorySize
	lastID  uint64
	subs    map[uint64]chan *model.Message
	nextSub uint64
}

// NewHub returns an empty Hub.
func NewHub() *Hub {
	return &Hub{subs: make(map[uint64]chan *model.Message)}
}

// Post validates, stores and fans out a message.
// Subscribers share the returned message. Do not mutate it.
func (h *Hub) Post(author, text string) (*model.Message, error) {
	author = strings.TrimSpace(author)
	text = strings.TrimSpace(text)
	if err := validate(author, MaxAuthorLen, ErrInvalidAuthor); err != nil {
		return nil, err
	}
	if err := validate(text, MaxTextLen, ErrInvalidText); err != nil {
		return nil, err
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.lastID++
	msg := &model.Message{
		ID:        strconv.FormatUint(h.lastID, 10),
		Author:    author,
		Text:      text,
		CreatedAt: time.Now().UTC(),
	}

	h.history[(h.lastID-1)%HistorySize] = msg

	for id, ch := range h.subs {
		select {
		case ch <- msg:
		default:
			// Evict a slow client instead of blocking the poster.
			delete(h.subs, id)
			close(ch)
		}
	}
	return msg, nil
}

// validate returns errLen unless s has 1 to maxLen runes.
func validate(s string, maxLen int, errLen error) error {
	n := 0
	for range s {
		if n++; n > maxLen {
			return errLen
		}
	}
	if n == 0 {
		return errLen
	}
	return nil
}

// Recent returns up to n of the newest messages, oldest first. n must not be negative.
func (h *Hub) Recent(n int) []*model.Message {
	h.mu.RLock()
	defer h.mu.RUnlock()

	k := min(n, int(min(h.lastID, HistorySize)))
	out := make([]*model.Message, k)
	start := h.lastID - uint64(k)
	for i := range k {
		out[i] = h.history[(start+uint64(i))%HistorySize]
	}
	return out
}

// Subscribe returns a channel that receives every new message until ctx is done or the subscriber is evicted.
// It returns ErrTooManySubscribers when MaxSubscribers subscriptions are open.
func (h *Hub) Subscribe(ctx context.Context) (<-chan *model.Message, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.subs) >= MaxSubscribers {
		return nil, ErrTooManySubscribers
	}
	ch := make(chan *model.Message, SubscriberBuffer)
	id := h.nextSub
	h.nextSub++
	h.subs[id] = ch
	// The callback runs in its own goroutine and may take h.mu.
	context.AfterFunc(ctx, func() { h.unsubscribe(id) })
	return ch, nil
}

// Subscribers returns the number of open subscriptions.
func (h *Hub) Subscribers() int {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return len(h.subs)
}

// unsubscribe removes and closes a subscriber. Whoever deletes the entry closes the channel.
func (h *Hub) unsubscribe(id uint64) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if ch, ok := h.subs[id]; ok {
		delete(h.subs, id)
		close(ch)
	}
}
