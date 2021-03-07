package users

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func (user Model) CreateValidation() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Name, validation.Required),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 25)),
	)
}

func (user Model) UpdateValidation() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, is.Email),
		validation.Field(&user.Name),
		validation.Field(&user.Password, validation.Length(6, 25)),
	)
}

func (login Login) LoginValidation() error {
	return validation.ValidateStruct(&login,
		validation.Field(&login.Email, validation.Required, is.Email),
		validation.Field(&login.Password, validation.Required),
	)
}
