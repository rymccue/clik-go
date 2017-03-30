package main

import (
	"crypto/rand"
	"io"
	"log"

	_ "github.com/lib/pq"

	"golang.org/x/crypto/scrypt"

	"encoding/base32"
	"encoding/base64"

	"github.com/jmoiron/sqlx"
)

const (
	SaltBytes = 32
	HashBytes = 64
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Open("postgres", "postgres://go:golang@localhost:5432/clik?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func generateSalt(password string) ([]byte, []byte) {
	salt := make([]byte, SaltBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}

	salt = []byte(base32.StdEncoding.EncodeToString(salt))[:SaltBytes]

	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, HashBytes)
	if err != nil {
		log.Fatal(err)
	}

	hash = []byte(base64.StdEncoding.EncodeToString(hash))[:HashBytes]

	return salt, hash
}

func DbCreateUser(user User, password string) User {
	salt, hash := generateSalt(password)
	userStmt := `insert into users (age, birthday, career, email, start_age, end_age, gender, info, first_name, last_name, looking_for, school, password, salt) values ($1, to_date($2, 'YYYY-MM-DD'), $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);`
	db.MustExec(userStmt, user.Age, user.Birthday, user.Career, user.Email, user.StartAge, user.EndAge, user.Gender, user.Info, user.FirstName, user.LastName, user.LookingFor, user.School, hash, salt)
	return user
}
