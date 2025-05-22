package utils

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

// GenerateShortCode generates a random 6-character string for use as a short code
func GenerateShortCode() (string, error) {
	// Generate 4 random bytes (will give us 6 characters in base64)
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Convert to base64 and remove any non-alphanumeric characters
	code := base64.URLEncoding.EncodeToString(b)
	code = strings.ReplaceAll(code, "-", "")
	code = strings.ReplaceAll(code, "_", "")

	// Take first 6 characters
	return code[:6], nil
}
