package memory

import (
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	repository := New(log.Logger)
	assert.Equal(t, log.Logger, repository.logger)
}
