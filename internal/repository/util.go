package repository

import (
	"crypto/sha256"
	"fmt"
)

func getObjectName(email string) string {
	emailBytes := []byte(email)
	emailHash := sha256.Sum256(emailBytes)
	return fmt.Sprintf("%x.png", emailHash)
}
