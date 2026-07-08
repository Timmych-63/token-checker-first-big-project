package service

import (
	"errors"

	"TOKENCHECKER/internal/model"
)

func RegisterUser(request model.RegisterRequest) error {
	if request.Login == "" {
		return errors.New("логин не может быть пустым")
	}

	if request.Password == "" {
		return errors.New("пароль не может быть пустым")
	}

	return nil
}
