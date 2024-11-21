package pkg

import "time"

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"string"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

type Category struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

type Expense struct {
	ID          string    `json:"id"`
	Amount      int64     `json:"amount"`
	Category    *Category `json:"category"`
	CreatedAt   time.Time `json:"date"`
	Description string    `json:"description"`
	User        *User     `json:"user"`
}

type Expenses []*Expense
