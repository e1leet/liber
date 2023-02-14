package config

import (
	"fmt"
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

type PostgresConfig struct {
	Username    string        `env:"USERNAME" env-required:"" env-description:"postgres username"`
	Password    string        `env:"PASSWORD" env-required:"" env-description:"postgres password"`
	Host        string        `env:"HOST" env-required:"" env-description:"postgres host"`
	Port        int           `env:"PORT" env-required:"" env-description:"postgres port"`
	Database    string        `env:"DATABASE" env-required:"" env-description:"postgres database"`
	SSLMode     string        `env:"SSL_MODE" env-required:"" env-description:"postgres ssl mode"`
	PingTimeout time.Duration `env:"PING_TIMEOUT" env-required:"" env-description:"postgres ping timeout"`
}

func (c PostgresConfig) URI() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.SSLMode,
	)
}

type Config struct {
	Server   ServerConfig   `env-prefix:"SERVER_"`
	Log      LogConfig      `env-prefix:"LOG_"`
	CORS     CORSConfig     `env-prefix:"CORS_"`
	Postgres PostgresConfig `env-prefix:"POSTGRES_"`
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
