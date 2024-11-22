package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) map[string]string {
	validate := validator.New()

	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	errors := make(map[string]string)
	for _, vErr := range validationErrors {
		errors[vErr.Field()] = fmt.Sprintf("failed on '%s' tag", vErr.Tag())
	}
	return errors
}
