package main

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"

	"golang.org/x/crypto/scrypt"

	"github.com/fatih/structs"
	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

const (
	SaltBytes        = 16
	EncodedSaltBytes = 32
	HashBytes        = 32
	EncodedHashBytes = 64
)

func init() {
	var err error
	db, err = sqlx.Open("postgres", "postgres://go:golang@localhost:5432/clik?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

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

func DbGetUser(id int) (*User, error) {
	user := &User{}
	tx := db.MustBegin()
	err := tx.Get(user, `select * from users where id = $1`, id)

	return user, err
}

func DbCreateUser(user *User) error {
	user.Salt = GenerateSalt()
	user.Password = HashPassword(user.Password, user.Salt)
	_, err := db.NamedExec(`
	insert into users
	(age, birthday, career, email, start_age, end_age, gender, info, first_name, last_name, looking_for, school, password, salt)
	values
	(:Age, to_date(:Birthday, 'YYYY-MM-DD'), :Career, :Email, :StartAge, :EndAge, :Gender, :Info, :FirstName, :LastName, :LookingFor, :School, :Password, :Salt)
	`, structs.Map(user))

	if err != nil {
		return err
	}

	tx := db.MustBegin()
	err = tx.Get(user, `select * from users where email = $1`, user.Email)

	return err
}
