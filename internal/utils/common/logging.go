package common

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func ConfigureLogging(level string) {
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(lvl).
		With().
		Timestamp().
		Caller().
		Logger()
}
