package util

import (
	"reflect"

	"example.com/expenses-tracker/internal/validators"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-playground/validator/v10"
)

const (
	requiredFieldMessage           = "This field is required"
	validEmailMessage              = "This field requires a valid email address"
	defaultFailedValidationMessage = "This field failed validation"
)

func FormatValidationMessages(obj any, fieldErrors validator.ValidationErrors) map[string]string {
	formatted := make(map[string]string, len(fieldErrors))

	for _, fieldError := range fieldErrors {
		tag := fieldError.Tag()
		field, ok := reflect.TypeOf(obj).FieldByName(fieldError.Field())
		fieldName := fieldError.Field()

		if ok {
			fieldName = string(field.Tag.Get("json"))
		}

		spew.Dump(fieldErrors)

		switch tag {
		case "required":
			formatted[fieldName] = requiredFieldMessage
		case "email":
			formatted[fieldName] = validEmailMessage
		case "validpassword":
			formatted[fieldName] = validators.ValidPasswordFieldMessage
		default:
			formatted[fieldName] = defaultFailedValidationMessage
		}
	}

	return formatted
}
