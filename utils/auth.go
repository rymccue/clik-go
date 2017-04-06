package utils

import (
	"bufio"
	"os"
	"time"

	"errors"

	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jeffmcnd/clik/models"
	"github.com/jeffmcnd/clik/repos"
)

const (
	PrivateKeyPath = "keys/clik"
	PublicKeyPath  = "keys/clik.pub"
)

// TODO(jeffmcnd): encrypt and decrypt JWT payloads

func getKey(path string) []byte {
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

func GenerateTokenForUser(user *models.User) (string, error) {
	key := getKey(PrivateKeyPath)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":     time.Now(),
		"exp":     time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC).Unix(),
		"email":   user.Email,
		"user_id": user.Id,
	})
	return token.SignedString(key)
}

func parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return getKey(PrivateKeyPath), nil
	})
}

func ValidateToken(tokenString string) error {
	token, err := parseToken(tokenString)
	if err != nil {
		return err
	}
	if token.Valid {
		return nil
	}
	return errors.New("something went wrong")
}

func GetUserIdFromToken(tokenString string) (int64, error) {
	token, err := parseToken(tokenString)
	if err != nil {
		return -1, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user, err := repos.DbGetUserByEmail(claims["email"].(string))
		return user.Id, err
	}
	return -1, nil
}
