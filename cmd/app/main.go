package main

import (
	"log"

	"TOKENCHECKER/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "./web/static")

	handlers.RegisterPageRouters(router)
	handlers.RegisterAPIRoutes(router)

	log.Println("Сервер запущен на 8080")

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
}
