package routers

import "github.com/gin-gonic/gin"

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == ""{

		}
	}
}