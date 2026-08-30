package handlers

import (
	"errors"
	"log"
	"net/http"

	"TOKENCHECKER/internal/model"
	"TOKENCHECKER/internal/service"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	messageService *service.MessageService
}

func RegisterMessageRoutes(
	router *gin.Engine,
	authService *service.AuthService,
	messageService *service.MessageService,
) {
	handler := &MessageHandler{
		messageService: messageService,
	}

	api := router.Group("/api")

	api.Use(
		requireAPIAuth(authService),
	)

	api.GET(
		"/message",
		handler.getMessage,
	)

	api.POST(
		"/message",
		handler.saveMessage,
	)
}

func (h *MessageHandler) getMessage(
	c *gin.Context,
) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			model.GetMessageResponse{},
		)
		return
	}

	text, err := h.messageService.GetMessage(userID)
	if err != nil {
		log.Printf(
			"Ошибка загрузки сообщения: %v",
			err,
		)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "внутренняя ошибка сервера",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		model.GetMessageResponse{
			Text: text,
		},
	)
}

func (h *MessageHandler) saveMessage(
	c *gin.Context,
) {
	userID, ok := getUserIDFromContext(c)
	if !ok {
		c.JSON(
			http.StatusInternalServerError,
			model.SaveMessageResponse{
				Message: "внутренняя ошибка сервера",
			},
		)
		return
	}

	var request model.SaveMessageRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			model.SaveMessageResponse{
				Message: "некорректные данные",
			},
		)
		return
	}

	err = h.messageService.SaveMessage(
		userID,
		request.Text,
	)
	if err != nil {
		if errors.Is(err, service.ErrMessageRequired) ||
			errors.Is(err, service.ErrMessageTooLong) {

			c.JSON(
				http.StatusBadRequest,
				model.SaveMessageResponse{
					Message: err.Error(),
				},
			)
			return
		}

		log.Printf(
			"Ошибка сохранения сообщения: %v",
			err,
		)

		c.JSON(
			http.StatusInternalServerError,
			model.SaveMessageResponse{
				Message: "внутренняя ошибка сервера",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		model.SaveMessageResponse{
			Message: "сообщение сохранено",
		},
	)
}

func getUserIDFromContext(
	c *gin.Context,
) (uint, bool) {
	value, exists := c.Get("userID")
	if !exists {
		return 0, false
	}

	userID, ok := value.(uint)
	if !ok {
		return 0, false
	}

	return userID, true
}
