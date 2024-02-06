package account

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func UsernameValidator(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile("^[a-z][a-z0-9]{3,15}$")
	return regex.MatchString(fl.Field().String())
}
