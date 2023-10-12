package validators

import (
	"errors"

	"github.com/SicParv1sMagna/mdhh_backend/internal/model"
)

func ValidateRegistrationData(user model.User) error {
	if user.FirstName == "" {
		return errors.New("имя должно быть заполнено")
	}

	if user.SecondName == "" {
		return errors.New("фамилия должна быть заполнена")
	}

	if user.Email == "" {
		return errors.New("поле email должно быть заполнено")
	}

	if user.Password == "" {
		return errors.New("поле пароль должно быть заполнено")
	}

	if user.Password != user.RepeatPassword {
		return errors.New("пароли должны совпадать")
	}

	if len(user.Password) > 20 {
		return errors.New("пароль должен быть не более 20 символов")
	}

	if len(user.Password) <= 8 {
		return errors.New("пароль должен быть больше 8 символов")
	}

	return nil
}

func ValidateAuthorizationData(user model.User) error {
	if !user.IsConfirmed {
		return errors.New("аккаунт не подтвержден")
	}

	if user.Email == "" || user.Password == "" {
		return errors.New("заполните все поля")
	}

	return nil
}
