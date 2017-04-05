package crypto

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"

	"golang.org/x/crypto/scrypt"
)

const (
	SaltBytes        = 16
	EncodedSaltBytes = 32
	HashBytes        = 32
	EncodedHashBytes = 64
)

// GenerateSalt generates a random 32 byte string
func GenerateSalt() string {
	saltBytes := make([]byte, SaltBytes)
	_, err := io.ReadFull(rand.Reader, saltBytes)
	if err != nil {
		log.Fatal("Failed to generate salt.")
	}
	salt := make([]byte, EncodedSaltBytes)
	hex.Encode(salt, saltBytes)
	return string(salt)
}

// HashPassword hashes a string with a given salt to a 64 byte string
func HashPassword(password string, salt string) string {
	hashBytes, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, HashBytes)
	if err != nil {
		log.Fatal("Failed to hash password.")
	}
	hash := make([]byte, EncodedHashBytes)
	hex.Encode(hash, hashBytes)
	return string(hash)
}
