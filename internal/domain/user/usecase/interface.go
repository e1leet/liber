package usecase

import (
	"context"

	"github.com/e1leet/liber/internal/domain/user/model"
)

type userRepository interface {
	Create(ctx context.Context, m map[string]interface{}) (*model.Storage, error)
	UserByID(ctx context.Context, id int) (*model.Storage, error)
	UserByEmail(ctx context.Context, email string) (*model.Storage, error)
	UserByUsername(ctx context.Context, username string) (*model.Storage, error)
}

type passwordManager interface {
	Hash(password string) string
}
