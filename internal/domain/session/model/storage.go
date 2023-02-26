package model

import (
	"time"

	sessionError "github.com/e1leet/liber/internal/domain/session/errors"
)

const (
	IDField        = "id"
	TokenField     = "token"
	UserIDField    = "user_id"
	ExpiresInField = "expires_in"
	CreatedAtField = "created_at"
)

type Storage struct {
	ID        int
	Token     string
	UserID    int
	ExpiresIn int64
	CreatedAt time.Time
}

func NewCreateMapper(token int, userID int, expiresIn int64) map[string]interface{} {
	return map[string]interface{}{
		TokenField:     token,
		UserIDField:    userID,
		ExpiresInField: expiresIn,
	}
}

func NewStorageFromCreateMapper(m map[string]interface{}) (*Storage, error) {
	token, ok := m[TokenField].(string)
	if !ok {
		return nil, sessionError.ErrTokenNotSpecified
	}

	userID, ok := m[UserIDField].(int)
	if !ok {
		return nil, sessionError.ErrUserIDNotSpecified
	}

	expiresIn, ok := m[ExpiresInField].(int64)
	if !ok {
		return nil, sessionError.ErrExpiresInNotSpecified
	}

	return &Storage{
		Token:     token,
		UserID:    userID,
		ExpiresIn: expiresIn,
	}, nil
}
