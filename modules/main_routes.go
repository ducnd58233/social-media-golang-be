package module

import (
	component "social-media-be/components"

	"github.com/gin-gonic/gin"
)

func MainRoute(router *gin.Engine, appCtx component.AppContext) {
	v1 := router.Group("/v1")
	{
		v1.GET("/")
	}
}