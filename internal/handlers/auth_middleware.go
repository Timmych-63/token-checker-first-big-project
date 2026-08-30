package handlers

import (
	"errors"
	"log"
	"net/http"

	"TOKENCHECKER/internal/service"

	"github.com/gin-gonic/gin"
)

const sessionCookieName = "session_token"

func requirePageAuth(
	authService *service.AuthService,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(sessionCookieName)

		if err != nil || token == "" {
			clearSessionCookie(c)

			c.Redirect(
				http.StatusFound,
				"/login",
			)

			c.Abort()
			return
		}

		userID, err := authService.AuthenticateSessionToken(token)
		if err != nil {
			if !errors.Is(err, service.ErrUnauthorized) {
				log.Printf(
					"Ошибка проверки сессии: %v",
					err,
				)
			}

			clearSessionCookie(c)

			c.Redirect(
				http.StatusFound,
				"/login",
			)

			c.Abort()
			return
		}

		c.Set(
			"userID",
			userID,
		)

		c.Next()
	}
}

func requireAPIAuth(
	authService *service.AuthService,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(sessionCookieName)

		if err != nil || token == "" {
			clearSessionCookie(c)

			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "пользователь не авторизован",
				},
			)
			return
		}

		userID, err := authService.AuthenticateSessionToken(token)
		if err != nil {
			if errors.Is(err, service.ErrUnauthorized) {
				clearSessionCookie(c)

				c.AbortWithStatusJSON(
					http.StatusUnauthorized,
					gin.H{
						"message": "пользователь не авторизован",
					},
				)
				return
			}

			log.Printf(
				"Ошибка проверки API-сессии: %v",
				err,
			)

			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "внутренняя ошибка сервера",
				},
			)
			return
		}

		c.Set(
			"userID",
			userID,
		)

		c.Next()
	}
}

func clearSessionCookie(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie(
		sessionCookieName,
		"",
		-1,
		"/",
		"",
		false,
		true,
	)
}
