package password

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewManager(t *testing.T) {
	var (
		hasher  = SHA256Hasher()
		key     = ""
		manager = NewManager(hasher, key)
	)

	assert.Equal(t, fmt.Sprintf("%v", hasher), fmt.Sprintf("%v", manager.hasher))
	assert.Equal(t, key, manager.key)
}

func TestManager_Hash(t *testing.T) {
	testcases := []struct {
		password string
		key      string
		expected string
	}{
		{
			password: "test",
			key:      "",
			expected: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		},
		{
			password: "some",
			key:      "thing",
			expected: "3fc9b689459d738f8c88a3a48aa9e33542016b7a4052e001aaa536fca74813cb",
		},
		{
			password: "sponge",
			key:      "bob",
			expected: "f0e2e750791171b0391b682ec35835bd6a5c3f7c8d1d0191451ec77b4d75f240",
		},
	}

	for _, tt := range testcases {
		t.Run(fmt.Sprintf("hash_%s", tt.password), func(t *testing.T) {
			manager := NewManager(SHA256Hasher(), tt.key)
			actual := manager.Hash(tt.password)

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestManager_Check(t *testing.T) {
	testcases := []struct {
		password string
		key      string
		hashed   string
		expected bool
	}{
		{
			password: "test",
			key:      "",
			hashed:   "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			expected: true,
		},
		{
			password: "some",
			key:      "thing",
			hashed:   "3fc9b689459d738f8c88a3a48aa9e33542016b7a4052e001aaa536fca74813cb",
			expected: true,
		},
		{
			password: "sponge",
			key:      "bob",
			hashed:   "f0e2e750791171b0391b682ec35835bd6a5c3f7c8d1d0191451ec77b4d75f240",
			expected: true,
		},
		{
			password: "amazing",
			key:      "amazing",
			hashed:   "1232d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			expected: false,
		},
		{
			password: "some",
			key:      "some",
			hashed:   "9865d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
			expected: false,
		},
	}

	for _, tt := range testcases {
		t.Run(fmt.Sprintf("hash_%s", tt.password), func(t *testing.T) {
			manager := NewManager(SHA256Hasher(), tt.key)
			actual := manager.Check(tt.password, tt.hashed)

			assert.Equal(t, tt.expected, actual)
		})
	}
}
