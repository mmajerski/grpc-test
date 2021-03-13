package models

import valid "github.com/go-playground/validator/v10"

var validator *valid.Validate

func init() {
	validator = valid.New()
}

// Validate validates tags on object
func Validate(t interface{}) error {
	return validator.Struct(t)
}
