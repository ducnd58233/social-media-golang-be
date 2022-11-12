package postgin

import (
	"net/http"
	"social-media-be/common"
	component "social-media-be/components"
	postbiz "social-media-be/modules/posts/biz"
	postmodel "social-media-be/modules/posts/model"
	poststorage "social-media-be/modules/posts/storage"

	"github.com/gin-gonic/gin"
)

func UpdatePostHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data postmodel.Post

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewUpdatePostBiz(store)

		if err := biz.UpdatePostById(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
