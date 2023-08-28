package validx

import (
	"github.com/go-playground/validator/v10"
)

func Struct(s any) error {
	validate := validator.New()
	return validate.Struct(s)
}
