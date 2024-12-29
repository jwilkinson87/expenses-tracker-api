package validators

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

const (
	requiredFieldMessage           = "This field is required"
	validEmailMessage              = "This field requires a valid email address"
	defaultFailedValidationMessage = "This field failed validation"
	eqPasswordMessage              = "Passwords do not match"
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

		switch tag {
		case "required":
			formatted[fieldName] = requiredFieldMessage
		case "email":
			formatted[fieldName] = validEmailMessage
		case "validpassword":
			formatted[fieldName] = ValidPasswordFieldMessage
		case "eqfield":
			// Get the name of the field to compare against
			paramField := fieldError.Param()
			formatted[fieldName] = fmt.Sprintf("This field must match %s", paramField)
		default:
			formatted[fieldName] = defaultFailedValidationMessage
		}
	}

	return formatted
}
