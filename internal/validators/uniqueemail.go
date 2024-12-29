package validators

import (
	"example.com/expenses-tracker/internal/repositories"
	"github.com/go-playground/validator/v10"
)

const (
	UniqueEmailFieldMessage = "This email address already has an account registered with it. Please try logging in with that, or initiate a reset password request"
)

func UniqueEmail(repo repositories.UserAuthRepository) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return true
	}
}
