package models

import "time"

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
