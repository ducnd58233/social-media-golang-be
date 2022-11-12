package postgin

import (
	"net/http"
	"social-media-be/common"
	component "social-media-be/components"
	postbiz "social-media-be/modules/posts/biz"
	poststorage "social-media-be/modules/posts/storage"

	"github.com/gin-gonic/gin"
)

func GetPostHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewGetPostBiz(store)

		data, err := biz.GetPost(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		data.Mask(false)
		data.MaskOwnerId(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
