package validator

import (
	"github.com/go-playground/validator"
)

var validate = validator.New()

func ValidateStruct(obj interface{}) error {
	return validate.Struct(obj)
}
