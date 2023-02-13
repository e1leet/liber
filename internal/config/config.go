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

type LogConfig struct {
	Level string `env:"LEVEL" env-required:"" env-description:"log level"`
}

type CORSConfig struct {
	AllowedOrigins   []string `env:"ALLOWED_ORIGINS" env-default:"" env-description:"CORS allowed origins"`
	AllowedMethods   []string `env:"ALLOWED_METHODS" env-default:"" env-description:"CORS allowed method"`
	AllowedHeaders   []string `env:"ALLOWED_HEADERS" env-default:"" env-description:"CORS allowed headers"`
	ExposedHeaders   []string `env:"EXPOSED_HEADERS" env-default:"" env-description:"CORS exposed origins"`
	AllowCredentials bool     `env:"ALLOW_CREDENTIALS" env-default:"false" env-description:"CORS allow credentials"`
	MaxAge           int      `env:"MAX_AGE" env-default:"5" env-description:"CORS request cache max age"`
}

type Config struct {
	Server ServerConfig `env-prefix:"SERVER_"`
	Log    LogConfig    `env-prefix:"LOG_"`
	CORS   CORSConfig   `env-prefix:"CORS_"`
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
