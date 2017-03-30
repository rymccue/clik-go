package main

type User struct {
	Id         int64  `json:"id"`
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
	CreatedAt  string `json:"-" db:"created_at"`
	UpdatedAt  string `json:"-" db:"updated_at"`
	Password   string `json:"-"`
	Salt       string `json:"-"`
}

type Users []User

type UserQueue []struct {
	User
	Image
}
