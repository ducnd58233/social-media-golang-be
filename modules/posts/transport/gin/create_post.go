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

func CreatePostHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data postmodel.Post

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.OwnerId = requester.GetUserId()

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewCreatePostBiz(store)

		if err := biz.CreatePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DbTypePost)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
