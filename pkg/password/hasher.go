package password

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hasher func(password string, key string) string

func SHA256Hasher() Hasher {
	return func(password string, key string) string {
		h := sha256.New()
		h.Write([]byte(password + key))

		return hex.EncodeToString(h.Sum(nil))
	}
}
