package validation

import (
	"context"

	"example.com/expenses-tracker/api/internal/repositories"
	"github.com/go-playground/validator/v10"
)

const (
	UniqueEmailFieldMessage = "This email address already has an account registered with it. Please try logging in with that, or initiate a reset password request"
)

func UniqueEmail(repo repositories.UserRepository) validator.Func {
	return func(fl validator.FieldLevel) bool {
		ctx := context.Background()
		email := fl.Field().String()
		user, err := repo.GetUserByEmailAddress(ctx, email)
		if err != nil {
			return true
		}

		return user == nil
	}
}
