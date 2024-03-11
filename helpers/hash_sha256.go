package helpers

import (
	"crypto/sha256"
	"fmt"
)

func HashSHA256(payload string) string {
	return fmt.Sprintf(
		"%x",
		sha256.Sum256(
			[]byte(payload),
		),
	)
}
