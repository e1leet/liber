package usecase

import (
	"context"

	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/stretchr/testify/mock"
)

type userRepositoryMock struct {
	mock.Mock
}

func (m *userRepositoryMock) Create(ctx context.Context, mapper map[string]interface{}) (*model.Storage, error) {
	args := m.Called(ctx, mapper)

	return args.Get(0).(*model.Storage), args.Error(1)
}

func (m *userRepositoryMock) UserByID(ctx context.Context, id int) (*model.Storage, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(*model.Storage), args.Error(1)
}

func (m *userRepositoryMock) UserByEmail(ctx context.Context, email string) (*model.Storage, error) {
	args := m.Called(ctx, email)

	return args.Get(0).(*model.Storage), args.Error(1)
}

func (m *userRepositoryMock) UserByUsername(ctx context.Context, username string) (*model.Storage, error) {
	args := m.Called(ctx, username)

	return args.Get(0).(*model.Storage), args.Error(1)
}

type passwordManagerMock struct {
	mock.Mock
}

func (m *passwordManagerMock) Hash(password string) string {
	args := m.Called(password)
	return args.String(0)
}
