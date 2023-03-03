package encrypt

import (
	"crypto/sha256"
	"fmt"
)

// EncryptWithSalt encrypt data with salt
func EncryptWithSalt(data, salt string) string {
	h1 := sha256.New()
	h1.Write([]byte(data))

	s1 := fmt.Sprintf("%x", h1.Sum(nil))

	h2 := sha256.New()
	h2.Write([]byte(s1 + salt))

	return fmt.Sprintf("%x", h2.Sum(nil))
}
