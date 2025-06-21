package middlewares

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func TokenAuthMiddleware(cfg config.AuthConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, api.NewErrorResponse("No token provided"))
			c.Abort()
			return
		}

		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

		jwt := utils.NewJwtUtils(&utils.JwtConfig{
			Secret: cfg.Secret,
		})

		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, api.NewErrorResponse(err.Error()))
			c.Abort()
			return
		}

		c.Set("auth_data", claims.Payload)

		c.Next()
	}
}
