package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(f1 validator.FieldLevel) bool {
	mobile := f1.Field().String()
	pattern := ""
	_, err := regexp.MatchString(pattern, mobile)
	if err != nil {
		return false
	}
	return false
}