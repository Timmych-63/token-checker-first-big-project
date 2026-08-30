package service

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"TOKENCHECKER/internal/model"
	"TOKENCHECKER/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

const SessionDuration = 7 * 24 * time.Hour

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

	ErrInvalidCredentials = errors.New(
		"неверный логин или пароль",
	)

	ErrUnauthorized = errors.New(
		"пользователь не авторизован",
	)
)

type AuthService struct {
	userRepository    *repository.UserRepository
	sessionRepository *repository.SessionRepository
}

func NewAuthService(
	userRepository *repository.UserRepository,
	sessionRepository *repository.SessionRepository,
) *AuthService {
	return &AuthService{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
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

func (s *AuthService) LoginUser(
	request model.LoginRequest,
) (string, error) {
	login := strings.TrimSpace(request.Login)
	password := request.Password

	if login == "" || password == "" {
		return "", ErrInvalidCredentials
	}

	user, err := s.userRepository.FindByLogin(login)
	if err != nil {
		return "", fmt.Errorf(
			"не удалось найти пользователя при входе: %w",
			err,
		)
	}

	if user == nil {
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		if errors.Is(
			err,
			bcrypt.ErrMismatchedHashAndPassword,
		) {
			return "", ErrInvalidCredentials
		}

		return "", fmt.Errorf(
			"не удалось проверить пароль: %w",
			err,
		)
	}

	token, tokenHash, err := generateSessionToken()
	if err != nil {
		return "", fmt.Errorf(
			"не удалось создать токен сессии: %w",
			err,
		)
	}

	session := &model.Session{
		UserID:    user.ID,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(SessionDuration),
	}

	err = s.sessionRepository.Create(session)
	if err != nil {
		return "", fmt.Errorf(
			"не удалось создать сессию: %w",
			err,
		)
	}

	return token, nil
}

func (s *AuthService) AuthenticateSessionToken(
	token string,
) (uint, error) {
	if strings.TrimSpace(token) == "" {
		return 0, ErrUnauthorized
	}

	tokenHash := hashSessionToken(token)

	session, err := s.sessionRepository.FindByTokenHash(
		tokenHash,
	)
	if err != nil {
		return 0, fmt.Errorf(
			"не удалось проверить сессию: %w",
			err,
		)
	}

	if session == nil {
		return 0, ErrUnauthorized
	}

	if !time.Now().Before(session.ExpiresAt) {
		err = s.sessionRepository.DeleteByTokenHash(
			tokenHash,
		)
		if err != nil {
			return 0, fmt.Errorf(
				"не удалось удалить просроченную сессию: %w",
				err,
			)
		}

		return 0, ErrUnauthorized
	}

	return session.UserID, nil
}

func (s *AuthService) LogoutUser(token string) error {
	if strings.TrimSpace(token) == "" {
		return nil
	}

	tokenHash := hashSessionToken(token)

	err := s.sessionRepository.DeleteByTokenHash(
		tokenHash,
	)
	if err != nil {
		return fmt.Errorf(
			"не удалось удалить сессию пользователя: %w",
			err,
		)
	}

	return nil
}

func generateSessionToken() (
	string,
	string,
	error,
) {
	randomBytes := make([]byte, 32)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", "", err
	}

	token := base64.RawURLEncoding.EncodeToString(
		randomBytes,
	)

	tokenHash := hashSessionToken(token)

	return token, tokenHash, nil
}

func hashSessionToken(token string) string {
	hash := sha256.Sum256(
		[]byte(token),
	)

	return hex.EncodeToString(
		hash[:],
	)
}
