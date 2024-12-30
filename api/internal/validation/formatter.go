package validation

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func FormatValidationMessages(obj any, fieldErrors validator.ValidationErrors) *map[string]string {
	messages := make(map[string]string)
	objType := reflect.TypeOf(obj)

	for _, fieldError := range fieldErrors {
		fieldName := fieldError.Field()
		tag := fieldError.Tag()

		// Find the struct field
		field, ok := objType.FieldByName(fieldName)
		if !ok {
			// If the field doesn't exist, use a default error message
			messages[fieldName] = "Invalid input"
			continue
		}

		// Use the JSON tag if available, otherwise default to the struct field name
		jsonTag := field.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = fieldName
		}

		// Extract the custom validation tag for the specific validation rule
		validationTag := fmt.Sprintf("validation[%s]", tag)
		if message, ok := field.Tag.Lookup(validationTag); ok {
			messages[jsonTag] = message
		} else {
			// Fallback message if no custom tag exists
			messages[jsonTag] = fmt.Sprintf("%s validation failed", jsonTag)
		}
	}

	return &messages
}
