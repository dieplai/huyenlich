// Package persistence — in-memory repo cho development, thay bằng Postgres ở production
package persistence

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/tarot/backend/internal/domain/reading"
)

type MemoryReadingRepo struct {
	mu   sync.RWMutex
	data map[uuid.UUID]*reading.Reading
}

func NewMemoryReadingRepo() *MemoryReadingRepo {
	return &MemoryReadingRepo{data: make(map[uuid.UUID]*reading.Reading)}
}

func (r *MemoryReadingRepo) Save(ctx context.Context, reading *reading.Reading) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[reading.ID] = reading
	return nil
}

func (r *MemoryReadingRepo) FindByID(ctx context.Context, id uuid.UUID) (*reading.Reading, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if v, ok := r.data[id]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("reading not found: %s", id)
}

func (r *MemoryReadingRepo) FindByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*reading.Reading, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*reading.Reading
	for _, v := range r.data {
		if v.UserID != nil && *v.UserID == userID {
			result = append(result, v)
		}
	}
	return result, nil
}

func (r *MemoryReadingRepo) Update(ctx context.Context, reading *reading.Reading) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[reading.ID] = reading
	return nil
}
