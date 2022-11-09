package module

import (
	component "social-media-be/components"
	middleware "social-media-be/middlewares"
	uploadgin "social-media-be/modules/upload/transport/gin"
	usergin "social-media-be/modules/user/transport/gin"

	"github.com/gin-gonic/gin"
)

func MainRoute(router *gin.Engine, appCtx component.AppContext) {
	v1 := router.Group("/v1")
	{
		v1.POST("/register", usergin.RegisterHandler(appCtx))
		v1.POST("/login", usergin.LoginHandler(appCtx))
		v1.POST("/profile", middleware.RequiredAuth(appCtx), usergin.GetProfileHandler(appCtx))
		v1.POST("/upload", uploadgin.UploadHandler(appCtx))
	}
}
