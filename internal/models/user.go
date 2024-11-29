package models

import "time"

type User struct {
	ID         string    `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"string"`
	Password   string    `json:"-"`
	ExpiryTime time.Time `json:"expiry_time"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserToken struct {
	Value      string
	User       *User
	ExpiryTime *time.Time
}

func (u *UserToken) IsTokenValid() bool {
	return u.User.ExpiryTime.Before(time.Now())
}
