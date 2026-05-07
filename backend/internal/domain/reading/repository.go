package reading

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Save(ctx context.Context, r *Reading) error
	FindByID(ctx context.Context, id uuid.UUID) (*Reading, error)
	FindByUserID(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*Reading, error)
	Update(ctx context.Context, r *Reading) error
}
