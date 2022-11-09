package uploadgin

import (
	"net/http"
	"social-media-be/common"
	_ "image/jpeg"
	_ "image/png"
	component "social-media-be/components"
	uploadbiz "social-media-be/modules/upload/biz"
	uploadstorage "social-media-be/modules/upload/storage"

	"github.com/gin-gonic/gin"
)

func UploadHandler(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		imgStore := uploadstorage.NewSQLStore(db)
		biz := uploadbiz.NewUploadBiz(appCtx.GetCloudProvider(), imgStore)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
