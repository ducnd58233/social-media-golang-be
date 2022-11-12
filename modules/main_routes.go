package module

import (
	component "social-media-be/components"
	middleware "social-media-be/middlewares"
	postgin "social-media-be/modules/posts/transport/gin"
	uploadgin "social-media-be/modules/upload/transport/gin"
	usergin "social-media-be/modules/user/transport/gin"

	"github.com/gin-gonic/gin"
)

func MainRoute(router *gin.Engine, appCtx component.AppContext) {
	v1 := router.Group("/v1")
	{
		v1.GET("", middleware.RequiredAuth(appCtx), postgin.ListPostHandler(appCtx))
		v1.POST("/register", usergin.RegisterHandler(appCtx))
		v1.POST("/login", usergin.LoginHandler(appCtx))
		v1.POST("/profile", middleware.RequiredAuth(appCtx), usergin.GetProfileHandler(appCtx))
		v1.POST("/upload", uploadgin.UploadHandler(appCtx))

		posts := v1.Group("/posts", middleware.RequiredAuth(appCtx))
		{
			posts.GET("/:id", postgin.GetPostHandler(appCtx))
			posts.GET("", postgin.ListPostHandler(appCtx))
			posts.GET("/user/:user_id", postgin.ListUserPostsHandler(appCtx))
			posts.POST("", postgin.CreatePostHandler(appCtx))
			posts.PATCH("/:id", postgin.UpdatePostHandler(appCtx))
			posts.DELETE("/:id", postgin.DeletePostHandler(appCtx))
		}
	}
}
