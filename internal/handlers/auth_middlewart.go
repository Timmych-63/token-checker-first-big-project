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
