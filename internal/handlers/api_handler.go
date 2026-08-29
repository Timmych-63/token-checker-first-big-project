package handlers

import (
	"errors"
	"log"
	"net/http"

	"TOKENCHECKER/internal/model"
	"TOKENCHECKER/internal/service"

	"github.com/gin-gonic/gin"
)

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
