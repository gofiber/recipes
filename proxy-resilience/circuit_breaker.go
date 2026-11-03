package main

import (
	"errors"
	"sync"
	"time"
)

// State represents the current circuit breaker operating state.
type State int

const (
	// StateClosed allows requests to pass through to the upstream.
	StateClosed State = iota
	// StateHalfOpen allows a limited trial request to check upstream health.
	StateHalfOpen
	// StateOpen fails fast without contacting the failing upstream.
	StateOpen
)

func (s State) String() string {
	switch s {
	case StateClosed:
		return "CLOSED"
	case StateHalfOpen:
		return "HALF-OPEN"
	case StateOpen:
		return "OPEN"
	default:
		return "UNKNOWN"
	}
}

// ErrCircuitOpen is returned when the circuit breaker is in Open state.
var ErrCircuitOpen = errors.New("circuit breaker is open: upstream service unavailable")

// CircuitBreakerConfig holds settings for failure thresholds and recovery.
type CircuitBreakerConfig struct {
	// MaxFailures is the consecutive failure threshold to trip the circuit open.
	MaxFailures int
	// Timeout is the cooldown duration before transitioning from Open to Half-Open.
	Timeout time.Duration
}

// CircuitBreaker tracks failure metrics and coordinates upstream availability.
type CircuitBreaker struct {
	mu          sync.RWMutex
	state       State
	failures    int
	maxFailures int
	timeout     time.Duration
	lastState   time.Time
}

// NewCircuitBreaker initializes a circuit breaker with the given configuration.
func NewCircuitBreaker(cfg CircuitBreakerConfig) *CircuitBreaker {
	if cfg.MaxFailures <= 0 {
		cfg.MaxFailures = 3
	}
	if cfg.Timeout <= 0 {
		cfg.Timeout = 5 * time.Second
	}
	return &CircuitBreaker{
		state:       StateClosed,
		maxFailures: cfg.MaxFailures,
		timeout:     cfg.Timeout,
		lastState:   time.Now(),
	}
}

// Execute runs the provided upstream call under circuit breaker protection.
func (cb *CircuitBreaker) Execute(fn func() error) error {
	if err := cb.beforeCall(); err != nil {
		return err
	}

	err := fn()
	cb.afterCall(err)
	return err
}

func (cb *CircuitBreaker) beforeCall() error {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	now := time.Now()
	if cb.state == StateOpen {
		if now.Sub(cb.lastState) > cb.timeout {
			cb.state = StateHalfOpen
			cb.lastState = now
			return nil
		}
		return ErrCircuitOpen
	}

	return nil
}

func (cb *CircuitBreaker) afterCall(err error) {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	now := time.Now()
	if err != nil {
		cb.failures++
		if cb.state == StateHalfOpen || cb.failures >= cb.maxFailures {
			cb.state = StateOpen
			cb.lastState = now
		}
		return
	}

	// Success resets failures and restores closed state
	cb.failures = 0
	cb.state = StateClosed
	cb.lastState = now
}

// State returns the current circuit state safely.
func (cb *CircuitBreaker) State() State {
	cb.mu.RLock()
	defer cb.mu.RUnlock()
	return cb.state
}
