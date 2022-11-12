package postgin

import (
	"net/http"
	"social-media-be/common"
	component "social-media-be/components"
	postbiz "social-media-be/modules/posts/biz"
	poststorage "social-media-be/modules/posts/storage"

	"github.com/gin-gonic/gin"
)

func ListUserPostsHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("user_id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewListPostsBiz(store)

		result, err := biz.ListPostByUserId(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
			result[i].MaskOwnerId(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, nil, nil))
	}
}
