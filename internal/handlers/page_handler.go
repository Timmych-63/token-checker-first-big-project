package handlers

import (
	"net/http"

	"TOKENCHECKER/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterPageRoutes(
	router *gin.Engine,
	authService *service.AuthService,
) {
	router.GET(
		"/",
		showHomePage,
	)

	router.GET(
		"/register",
		showRegisterPage,
	)

	router.GET(
		"/login",
		showLoginPage,
	)

	router.GET(
		"/cabinet",
		requirePageAuth(authService),
		showCabinetPage,
	)
}

func showHomePage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "Главная страница",
		},
	)
}

func showRegisterPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"register.html",
		gin.H{
			"title": "Регистрация",
		},
	)
}

func showLoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Вход",
		},
	)
}

func showCabinetPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"cabinet.html",
		gin.H{
			"title": "Личный кабинет",
		},
	)
}
