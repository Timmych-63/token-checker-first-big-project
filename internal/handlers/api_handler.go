package handlers

import (
	"net/http"

	"TOKENCHECKER/internal/model"
	"TOKENCHECKER/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.POST("/register", registerUser)
}

func registerUser(c *gin.Context) {
	var request model.RegisterRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "некорректные данные",
		})
		return
	}

	err = service.RegisterUser(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.RegisterResponse{
		Message: "регистрация выполнена",
	})
}
