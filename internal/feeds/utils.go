package feeds

import (
	"crypto/sha1"
	"encoding/hex"
)

func generateID(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
