package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := Authenticate(c.GetHeader("Authorization"))
		if err != nil {
			failedResponse(c)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func failedResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
		"code":    403,
		"message": "去登陆",
	})
}
