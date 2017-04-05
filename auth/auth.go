package auth

import (
	"bufio"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	PrivateKeyPath = "keys/clik"
	PublicKeyPath  = "keys/clik.pub"
)

func GetKey(path string) []byte {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	info, _ := file.Stat()
	size := info.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(pembytes)

	return pembytes
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return GetKey(PublicKeyPath), nil
	})

	if err != nil {
		return false
	}

	if token.Valid {
		return true
	}
	return false
}
