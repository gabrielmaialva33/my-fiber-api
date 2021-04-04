package validators

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go-api/src/app/modules/user/models"
)

func CreateUserValidator(user models.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required, validation.Length(4, 100)),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Username, validation.Length(4, 50)),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 100)),
	)
}

func LoginUserValidator(user models.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required, validation.Length(6, 100)),
	)
}
