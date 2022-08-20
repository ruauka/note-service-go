// Package validate Package validate
package validate

import (
	"github.com/go-playground/validator/v10"

	"web/internal/utils"
)

// InputJSONValidate input JSON validation.
func InputJSONValidate(inputJSON interface{}) error {
	err := utils.Validate.Struct(inputJSON)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}
