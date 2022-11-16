package middleware

import (
	"net/http"
	"akiba/utils/token"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc{
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized,"Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}