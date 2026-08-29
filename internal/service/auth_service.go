package service

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"

	"TOKENCHECKER/internal/model"
	"TOKENCHECKER/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrLoginRequired = errors.New(
		"логин не может быть пустым",
	)

	ErrLoginTooShort = errors.New(
		"логин должен содержать минимум 3 символа",
	)

	ErrLoginTooLong = errors.New(
		"логин должен содержать максимум 50 символов",
	)

	ErrPasswordRequired = errors.New(
		"пароль не может быть пустым",
	)

	ErrPasswordTooShort = errors.New(
		"пароль должен содержать минимум 8 символов",
	)

	ErrPasswordTooLong = errors.New(
		"пароль слишком длинный",
	)

	ErrLoginAlreadyExists = errors.New(
		"пользователь с таким логином уже существует",
	)
)

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(
	userRepository *repository.UserRepository,
) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) RegisterUser(
	request model.RegisterRequest,
) error {
	login := strings.TrimSpace(request.Login)
	password := request.Password

	if login == "" {
		return ErrLoginRequired
	}

	loginLength := utf8.RuneCountInString(login)

	if loginLength < 3 {
		return ErrLoginTooShort
	}

	if loginLength > 50 {
		return ErrLoginTooLong
	}

	if strings.TrimSpace(password) == "" {
		return ErrPasswordRequired
	}

	if utf8.RuneCountInString(password) < 8 {
		return ErrPasswordTooShort
	}

	if len([]byte(password)) > 72 {
		return ErrPasswordTooLong
	}

	existingUser, err := s.userRepository.FindByLogin(login)
	if err != nil {
		return fmt.Errorf(
			"не удалось проверить существование пользователя: %w",
			err,
		)
	}

	if existingUser != nil {
		return ErrLoginAlreadyExists
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fmt.Errorf(
			"не удалось захешировать пароль: %w",
			err,
		)
	}

	user := &model.User{
		Login:        login,
		PasswordHash: string(passwordHash),
	}

	err = s.userRepository.Create(user)
	if err != nil {
		return fmt.Errorf(
			"не удалось создать пользователя: %w",
			err,
		)
	}

	return nil
}
