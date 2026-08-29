package handlers

import (
	"errors"
	"log"
	"net/http"

	"TOKENCHECKER/internal/model"
	"TOKENCHECKER/internal/service"

	"github.com/gin-gonic/gin"
)

const sessionCookieMaxAge = 7 * 24 * 60 * 60

type APIHandler struct {
	authService *service.AuthService
}

func RegisterAPIRoutes(
	router *gin.Engine,
	authService *service.AuthService,
) {
	handler := &APIHandler{
		authService: authService,
	}

	api := router.Group("/api")

	api.POST("/register", handler.registerUser)
	api.POST("/login", handler.loginUser)
}

func (h *APIHandler) registerUser(c *gin.Context) {
	var request model.RegisterRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.RegisterResponse{
				Message: "некорректные данные",
			},
		)
		return
	}

	err = h.authService.RegisterUser(request)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrLoginRequired),
			errors.Is(err, service.ErrLoginTooShort),
			errors.Is(err, service.ErrLoginTooLong),
			errors.Is(err, service.ErrPasswordRequired),
			errors.Is(err, service.ErrPasswordTooShort),
			errors.Is(err, service.ErrPasswordTooLong):

			c.JSON(
				http.StatusBadRequest,
				model.RegisterResponse{
					Message: err.Error(),
				},
			)

		case errors.Is(err, service.ErrLoginAlreadyExists):
			c.JSON(
				http.StatusConflict,
				model.RegisterResponse{
					Message: err.Error(),
				},
			)

		default:
			log.Printf(
				"Ошибка регистрации пользователя: %v",
				err,
			)

			c.JSON(
				http.StatusInternalServerError,
				model.RegisterResponse{
					Message: "внутренняя ошибка сервера",
				},
			)
		}

		return
	}

	c.JSON(
		http.StatusCreated,
		model.RegisterResponse{
			Message: "регистрация выполнена",
		},
	)
}

func (h *APIHandler) loginUser(c *gin.Context) {
	var request model.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.LoginResponse{
				Message: "некорректные данные",
			},
		)
		return
	}

	token, err := h.authService.LoginUser(request)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			c.JSON(
				http.StatusUnauthorized,
				model.LoginResponse{
					Message: err.Error(),
				},
			)
			return
		}

		log.Printf(
			"Ошибка входа пользователя: %v",
			err,
		)

		c.JSON(
			http.StatusInternalServerError,
			model.LoginResponse{
				Message: "внутренняя ошибка сервера",
			},
		)
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie(
		"session_token",
		token,
		sessionCookieMaxAge,
		"/",
		"",
		false,
		true,
	)

	c.JSON(
		http.StatusOK,
		model.LoginResponse{
			Message: "вход выполнен",
		},
	)
}
