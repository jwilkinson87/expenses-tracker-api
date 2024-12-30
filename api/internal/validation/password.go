package validation

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

const (
	ValidPasswordFieldMessage = "This field requires a valid password. Please enter a password that has at least 1 upper character, 1 lower character, 1 number, 1 special character, and is at least 7 characters in length"
)

var ValidPassword validator.Func = func(fl validator.FieldLevel) bool {
	s := fl.Field().Interface().(string)

	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
