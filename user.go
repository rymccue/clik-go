package main

type User struct {
	Id         int    `json:"id"`
	Age        int    `json:"age"`
	Birthday   string `json:"birthday"`
	Career     string `json:"career"`
	Email      string `json:"email"`
	StartAge   int    `json:"start_age" db:"start_age"`
	EndAge     int    `json:"end_age" db:"end_age"`
	Gender     string `json:"gender"`
	Info       string `json:"info"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	LookingFor string `json:"looking_for" db:"looking_for"`
	School     string `json:"school"`
	CreatedAt  string `json:"created_at" db:"created_at"`
	UpdatedAt  string `json:"updated_at" db:"updated_at"`
}

type Users []User
