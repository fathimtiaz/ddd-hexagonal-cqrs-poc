package input

import "github.com/go-playground/validator/v10"

var (
	v = validator.New()
)

func ValidatePayload(payload any) error {
	return v.Struct(payload)
}
