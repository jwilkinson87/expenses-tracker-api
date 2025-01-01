package models

import (
	"time"

	"example.com/expenses-tracker/pkg/requests"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email_address"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type ResetToken struct {
	ID         string
	CreatedAt  time.Time
	ExpiryTime time.Time
	ResetToken string
	User       *User
}

func (u *User) FromUserRequest(request *requests.CreateUserRequest) error {
	id, _ := uuid.NewV7()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.ID = id.String()
	u.FirstName = request.FirstName
	u.LastName = request.LastName
	u.Email = request.EmailAddress
	u.Password = string(hashedPassword)
	u.CreatedAt = time.Now()

	return nil
}
