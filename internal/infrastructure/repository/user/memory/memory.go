package memory

import (
	"sync"

	"github.com/e1leet/liber/internal/domain/user/model"
	"github.com/rs/zerolog"
)

type Memory struct {
	users  map[int]*model.Storage
	mu     sync.Mutex
	logger zerolog.Logger
}

func New(logger zerolog.Logger) *Memory {
	return &Memory{
		users:  make(map[int]*model.Storage),
		logger: logger,
	}
}
