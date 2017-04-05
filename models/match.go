package models

type Match struct {
	Id        int64  `json:"id"`
	Age       int    `json:"age"`
	Career    string `json:"career"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	School    string `json:"school"`
	MatchId   int64  `json:"match_id" db:"match_id"`
	ImageUrl
}

type Matches []Match
