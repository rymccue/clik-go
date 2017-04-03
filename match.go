package main

type Match struct {
	Id        int64  `json:"id"`
	Age       int    `json:"age"`
	Career    string `json:"career"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	School    string `json:"school"`
	ImageUrl
}

type Matches []Match
