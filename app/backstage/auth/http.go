package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpNeedLogin(c *gin.Context) {
	user, exist := c.Get("user")
	c.JSON(http.StatusOK, map[string]interface{}{
		"user":  user,
		"exist": exist,
	})
}
