package util

import (
	"reflect"

	"github.com/go-playground/validator/v10"
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
			formatted[fieldName] = "this field is required"
		case "email":
			formatted[fieldName] = "this field requires a valid email address"
		default:
			formatted[fieldName] = "this field failed validation"
		}
	}

	return formatted
}
