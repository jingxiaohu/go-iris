package support

import "gopkg.in/go-playground/validator.v9"

var G_validate *validator.Validate
func InitValidator() {
	G_validate = validator.New()
}
