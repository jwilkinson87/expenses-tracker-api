package models

import (
	"time"

	"example.com/expenses-tracker/internal/requests"
	"github.com/hashicorp/go-uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"string"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type UserToken struct {
	ID         string
	Value      string
	User       *User
	ExpiryTime *time.Time
}

type ResetToken struct {
	ID         string
	CreatedAt  time.Time
	ExpiryTime time.Time
	ResetToken string
	User       *User
}

func (u *User) FromUserRequest(request *requests.CreateUserRequest) error {
	id, _ := uuid.GenerateUUID()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.ID = id
	u.FirstName = request.FirstName
	u.LastName = request.LastName
	u.Email = request.EmailAddress
	u.Password = string(hashedPassword)
	u.CreatedAt = time.Now()

	return nil
}

func (u *UserToken) IsTokenValid() bool {
	return u.ExpiryTime.Before(time.Now())
}
