package main

import (
	_ "github.com/lib/pq"

	"github.com/jeffmcnd/clik/repos"
	"github.com/jeffmcnd/clik/web"
)

func main() {
	// databaseString := "user=go password=golang dbname=clik sslmode=disable"
	databaseString := "postgres://go:golang@localhost:5432/clik?sslmode=disable"
	repos.NewDatabaseConnection(databaseString)

	s := web.CreateServer()
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
