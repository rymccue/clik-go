package repos

import (
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func NewDatabaseConnection(url string) {
	db, _ = sqlx.Open("postgres", url)
}
