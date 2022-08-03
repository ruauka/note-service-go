package validate

import (
	"github.com/go-playground/validator/v10"

	"web/internal/utils"
)

// InputJsonValidate - валидация входящей структуры.
func InputJsonValidate(inputJson interface{}) error {
	err := utils.Validate.Struct(inputJson)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}
