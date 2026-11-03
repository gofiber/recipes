package chat

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"testing"
	"testing/synctest"

	"github.com/gofiber/recipes/gqlgen/graph/model"
)

func mustPost(t *testing.T, h *Hub, text string) *model.Message {
	t.Helper()
	msg, err := h.Post("ana", text)
	if err != nil {
		t.Fatalf("Post(%q): %v", text, err)
	}
	return msg
}

func mustSubscribe(t *testing.T, h *Hub, ctx context.Context) <-chan *model.Message {
	t.Helper()
	ch, err := h.Subscribe(ctx)
	if err != nil {
		t.Fatalf("Subscribe: %v", err)
	}
	return ch
}

func ids(msgs []*model.Message) []string {
	out := make([]string, len(msgs))
	for i, m := range msgs {
		out[i] = m.ID
	}
	return out
}

func idRange(from, to int) []string {
	out := make([]string, 0, to-from+1)
	for i := from; i <= to; i++ {
		out = append(out, strconv.Itoa(i))
	}
	return out
}

func TestRecent(t *testing.T) {
	h := NewHub()
	if got := h.Recent(20); got == nil || len(got) != 0 {
		t.Fatalf("Recent on empty hub = %v, want empty non-nil slice", got)
	}

	for i := range 3 {
		mustPost(t, h, fmt.Sprint(i))
	}
	if got, want := ids(h.Recent(20)), idRange(1, 3); !slices.Equal(got, want) {
		t.Fatalf("Recent(20) = %v, want %v", got, want)
	}
	if got := h.Recent(0); got == nil || len(got) != 0 {
		t.Fatalf("Recent(0) = %v, want empty non-nil slice", got)
	}
}

func TestRecentRingWraparound(t *testing.T) {
	h := NewHub()
	for i := range HistorySize + 1 {
		mustPost(t, h, fmt.Sprint(i))
	}
	if got, want := ids(h.Recent(HistorySize)), idRange(2, HistorySize+1); !slices.Equal(got, want) {
		t.Fatalf("Recent(%d) = %v, want %v", HistorySize, got, want)
	}
	if got, want := ids(h.Recent(5)), idRange(HistorySize-3, HistorySize+1); !slices.Equal(got, want) {
		t.Fatalf("Recent(5) = %v, want %v", got, want)
	}
	if got := len(h.Recent(HistorySize * 2)); got != HistorySize {
		t.Fatalf("len(Recent(%d)) = %d, want %d", HistorySize*2, got, HistorySize)
	}

	// Wrap more than once and from a non-zero head.
	for i := range 2*HistorySize + 37 {
		mustPost(t, h, fmt.Sprint(i))
	}
	last := 3*HistorySize + 38
	if got, want := ids(h.Recent(HistorySize)), idRange(last-HistorySize+1, last); !slices.Equal(got, want) {
		t.Fatalf("after wrap Recent(%d) = %v, want %v", HistorySize, got, want)
	}
}

func TestPostValidation(t *testing.T) {
	tests := []struct {
		name, author, text string
		want               error
	}{
		{"empty author", "", "hi", ErrInvalidAuthor},
		{"blank author", "   ", "hi", ErrInvalidAuthor},
		{"long author", strings.Repeat("a", MaxAuthorLen+1), "hi", ErrInvalidAuthor},
		{"empty text", "ana", "", ErrInvalidText},
		{"blank text", "ana", " \t\n", ErrInvalidText},
		{"long text", "ana", strings.Repeat("x", MaxTextLen+1), ErrInvalidText},
		{"max multibyte", strings.Repeat("ä", MaxAuthorLen), strings.Repeat("ü", MaxTextLen), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHub()
			msg, err := h.Post(tt.author, tt.text)
			if !errors.Is(err, tt.want) {
				t.Fatalf("err = %v, want %v", err, tt.want)
			}
			if tt.want != nil && (msg != nil || len(h.Recent(1)) != 0) {
				t.Fatal("invalid message was stored")
			}
		})
	}

	h := NewHub()
	msg := mustPost(t, h, "  hello  ")
	if msg.Text != "hello" || msg.Author != "ana" {
		t.Fatalf("message not trimmed: %+v", msg)
	}
}

func TestSlowSubscriberEvicted(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := NewHub()
		ctx, cancel := context.WithCancel(context.Background())
		slow := mustSubscribe(t, h, ctx)
		fast := mustSubscribe(t, h, context.Background())

		for i := range SubscriberBuffer {
			mustPost(t, h, fmt.Sprint(i))
			<-fast
		}
		if got := h.Subscribers(); got != 2 {
			t.Fatalf("subscribers = %d, want 2 before the buffer overflows", got)
		}

		mustPost(t, h, "overflow")
		<-fast
		if got := h.Subscribers(); got != 1 {
			t.Fatalf("subscribers = %d, want 1 after eviction", got)
		}

		// The buffered messages are still delivered, then the channel is closed.
		n := 0
		for range slow {
			n++
		}
		if n != SubscriberBuffer {
			t.Fatalf("slow subscriber drained %d messages, want %d", n, SubscriberBuffer)
		}

		// Cancel after eviction must not close the channel a second time.
		cancel()
		synctest.Wait()
		if got := h.Subscribers(); got != 1 {
			t.Fatalf("subscribers = %d, want 1", got)
		}
	})
}

func TestUnsubscribeOnCancel(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := NewHub()
		ctx, cancel := context.WithCancel(context.Background())
		ch := mustSubscribe(t, h, ctx)
		mustPost(t, h, "before")

		cancel()
		synctest.Wait()

		if got := h.Subscribers(); got != 0 {
			t.Fatalf("subscribers = %d, want 0", got)
		}
		if msg, ok := <-ch; !ok || msg.Text != "before" {
			t.Fatalf("buffered message lost: %v %v", msg, ok)
		}
		if _, ok := <-ch; ok {
			t.Fatal("channel still open after cancel")
		}

		mustPost(t, h, "after")
	})
}

func TestSubscribeCanceledContext(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := NewHub()
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		ch := mustSubscribe(t, h, ctx)
		synctest.Wait()
		if _, ok := <-ch; ok {
			t.Fatal("channel open for a canceled context")
		}
		if got := h.Subscribers(); got != 0 {
			t.Fatalf("subscribers = %d, want 0", got)
		}
	})
}

func TestMaxSubscribers(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := NewHub()
		ctx, cancel := context.WithCancel(context.Background())
		mustSubscribe(t, h, ctx)
		for range MaxSubscribers - 1 {
			mustSubscribe(t, h, t.Context())
		}
		if ch, err := h.Subscribe(t.Context()); !errors.Is(err, ErrTooManySubscribers) || ch != nil {
			t.Fatalf("Subscribe over the cap = %v, %v, want nil, %v", ch, err, ErrTooManySubscribers)
		}
		if got := h.Subscribers(); got != MaxSubscribers {
			t.Fatalf("subscribers = %d, want %d", got, MaxSubscribers)
		}

		// A canceled subscription frees its slot.
		cancel()
		synctest.Wait()
		mustSubscribe(t, h, t.Context())
	})
}
