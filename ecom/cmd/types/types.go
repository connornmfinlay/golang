package types

import "time"

type User struct {
	ID        int    		`json:"id"`
	FirstName string 		`json:"firstname"`
	LastName  string 		`json:"lastname"`
	Email     string  	`json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstame"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
}
