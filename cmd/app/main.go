package main

import (
	"log"

	"TOKENCHECKER/internal/config"
	"TOKENCHECKER/internal/handlers"
	"TOKENCHECKER/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Ошибка конфигурации: ", err)
	}

	db, err := repository.ConnectDatabase(cfg)
	if err != nil {
		log.Fatal("Ошибка подключения к базе: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(
			"Ошибка получения подключения для закрытия: ",
			err,
		)
	}

	defer func() {
		err := sqlDB.Close()
		if err != nil {
			log.Println(
				"Ошибка закрытия подключения к базе:",
				err,
			)
		}
	}()

	log.Printf(
		"Настройки базы загружены: host=%s port=%s db=%s user=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBUser,
	)

	router := gin.Default()

	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "./web/static")

	handlers.RegisterPageRouters(router)
	handlers.RegisterAPIRoutes(router)

	address := ":" + cfg.AppPort

	log.Printf(
		"Сервер запущен: http://localhost:%s",
		cfg.AppPort,
	)

	err = router.Run(address)
	if err != nil {
		log.Fatal("Ошибка при запуске сервера: ", err)
	}
}
