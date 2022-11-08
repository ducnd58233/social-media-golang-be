package usergin

import (
	"net/http"
	"social-media-be/common"
	component "social-media-be/components"

	"github.com/gin-gonic/gin"
)

func GetProfileHandler(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}