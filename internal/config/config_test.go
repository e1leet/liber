package config

import (
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
		}

		actual, err := New(testdataFolder + "config.env")
		require.NoError(t, err)

		assert.EqualValues(t, expected, actual)
	})
}
