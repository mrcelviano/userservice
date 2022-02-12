package validator

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"strings"
)

func ValidateEmail(email string) (bool, error) {
	email = strings.ToLower(strings.ReplaceAll(email, "%40", "@"))
	err := validation.Validate(email, validation.Required, is.Email)
	if err != nil {
		return false, err
	}
	return !strings.Contains(email, "|") && email != "", nil
}
