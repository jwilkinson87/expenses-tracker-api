package models

import "time"

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"string"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
