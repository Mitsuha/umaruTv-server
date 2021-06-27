package route

import (
	"github.com/gin-gonic/gin"
	appauth "umarutv/app/auth"
	bauth "umarutv/app/backstage/auth"
	"umarutv/common/auth"
)

func RegisterApiRoute(router *gin.RouterGroup) {
	bsRouter := router.Group("/backstage", auth.Middleware())
	{
		bsRouter.POST("/needLogin", bauth.HttpNeedLogin)
		bsRouter.POST("/animate/", )
	}

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", appauth.HttpLogin)
		authRouter.POST("/register", appauth.HttpRegister)
	}
}
