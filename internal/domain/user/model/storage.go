package model

import (
	"time"

	userError "github.com/e1leet/liber/internal/domain/user/errors"
)

const (
	IDField        = "id"
	EmailField     = "email"
	UsernameField  = "username"
	PasswordField  = "password"
	CreatedAtField = "created_at"
	UpdatedAtField = "updated_at"
)

type Storage struct {
	ID        int
	Email     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Storage) ToDomain() *User {
	return &User{
		ID:        s.ID,
		Email:     s.Email,
		Username:  s.Username,
		Password:  s.Password,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

func NewCreateMapper(email string, username string, password string) map[string]interface{} {
	return map[string]interface{}{
		EmailField:    email,
		UsernameField: username,
		PasswordField: password,
	}
}

func NewStorageFromCreateMapper(m map[string]interface{}) (*Storage, error) {
	email, ok := m[EmailField].(string)
	if !ok {
		return nil, userError.ErrEmailNotSpecified
	}

	username, ok := m[UsernameField].(string)
	if !ok {
		return nil, userError.ErrUsernameNotSpecified
	}

	password, ok := m[PasswordField].(string)
	if !ok {
		return nil, userError.ErrPasswordNotSpecified
	}

	return &Storage{
		Email:    email,
		Username: username,
		Password: password,
	}, nil
}
