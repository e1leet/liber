package password

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256Hasher(t *testing.T) {
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
			hasher := SHA256Hasher()
			actual := hasher(tt.password, tt.key)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
