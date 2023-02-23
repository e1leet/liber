package usecase

import (
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	repository := &userRepositoryMock{}
	manager := &passwordManagerMock{}

	usecase := New(repository, manager, log.Logger)

	assert.Equal(t, repository, usecase.userRepository)
	assert.Equal(t, manager, usecase.passwordManager)
	assert.Equal(t, log.Logger, usecase.logger)
}
