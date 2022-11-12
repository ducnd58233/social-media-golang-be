package postgin

import (
	"net/http"
	"social-media-be/common"
	component "social-media-be/components"
	postbiz "social-media-be/modules/posts/biz"
	poststorage "social-media-be/modules/posts/storage"

	"github.com/gin-gonic/gin"
)

func ListPostHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewListPostsBiz(store)

		result, err := biz.ListPost(c.Request.Context())

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
