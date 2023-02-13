package config

import (
	"log"
	"time"

	"github.com/e1leet/liber/pkg/errors"
	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	Host            string        `env:"HOST" env-required:"" env-description:"server host"`
	Port            int           `env:"PORT" env-required:"" env-description:"server port"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" env-required:"" env-description:"graceful shutdown timeout"`
}

type Config struct {
	Server ServerConfig `env-prefix:"SERVER_"`
}

func New(path string) (*Config, error) {
	log.Printf("gather config from %s", path)

	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		log.Println(err)
		return nil, errors.Wrap(err, "failed to gather config")
	}

	return cfg, nil
}
