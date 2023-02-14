package config

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testdataFolder = "testdata/"

func TestNew(t *testing.T) {
	t.Run("file_doesnt_exist", func(t *testing.T) {
		_, err := New("test.env")
		require.ErrorIs(t, err, os.ErrNotExist)
	})

	t.Run("gather_config", func(t *testing.T) {
		expected := &Config{
			Server: ServerConfig{
				Host:            "localhost",
				Port:            8000,
				ShutdownTimeout: 10 * time.Second,
			},
			Log: LogConfig{Level: "info"},
			CORS: CORSConfig{
				AllowedOrigins:   []string{"http://localhost:8000"},
				AllowedMethods:   []string{"GET"},
				AllowedHeaders:   []string{"Accept"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: true,
				MaxAge:           300,
			},
			Postgres: PostgresConfig{
				Username:    "postgres",
				Password:    "postgres",
				Host:        "localhost",
				Port:        5432,
				Database:    "postgres",
				SSLMode:     "disable",
				PingTimeout: time.Second * 5,
			},
		}

		actual, err := New(testdataFolder + "config.env")
		require.NoError(t, err)

		assert.EqualValues(t, expected, actual)
	})
}

func TestPostgresConfig_URI(t *testing.T) {
	cfg := PostgresConfig{
		Username:    "postgres",
		Password:    "postgres",
		Host:        "postgres",
		Port:        5432,
		Database:    "postgres",
		SSLMode:     "disable",
		PingTimeout: time.Second * 5,
	}

	expected := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSLMode,
	)
	assert.Equal(t, expected, cfg.URI())
}
