package usecase

import "github.com/rs/zerolog"

type Usecase struct {
	userRepository  userRepository
	logger          zerolog.Logger
	passwordManager passwordManager
}

func New(repository userRepository, logger zerolog.Logger, manager passwordManager) *Usecase {
	return &Usecase{
		userRepository:  repository,
		logger:          logger,
		passwordManager: manager,
	}
}
