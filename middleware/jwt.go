package middleware

import (
	"net/http"
	"technical-test/dto"
	"technical-test/helper"

	"github.com/gin-gonic/gin"
)

func (m *MiddlewareImpl) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helper.TokenValid(c.Request, m.Env)
		if err != nil {
			c.JSON(http.StatusUnauthorized, dto.ApiResponse{
				StatusCode: http.StatusUnauthorized,
				Success:    false,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
