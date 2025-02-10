package repository

import (
	"email-verification/domain"
	"fmt"
	"sync"
)

type MemoryRepository struct {
	verifications map[string]domain.Verification
	mu            sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		verifications: make(map[string]domain.Verification),
	}
}

func (r *MemoryRepository) Store(email string, v domain.Verification) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.verifications[email] = v
	return nil
}

func (r *MemoryRepository) Get(email string) (domain.Verification, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	v, exists := r.verifications[email]
	if !exists {
		return domain.Verification{}, fmt.Errorf("verification not found")
	}
	return v, nil
}

func (r *MemoryRepository) Delete(email string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.verifications, email)
	return nil
}
