package usergin

import (
	"net/http"
	"social-media-be/common"
	component "social-media-be/components"
	"social-media-be/components/tokenprovider/hasher"
	userbiz "social-media-be/modules/user/biz"
	usermodel "social-media-be/modules/user/model"
	userstorage "social-media-be/modules/user/storage"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var data usermodel.UserCreate

		db := appCtx.GetMainDBConnection()

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userStore := userstorage.NewSQLStore(db)
		bcrypt := hasher.NewBcryptHash()
		biz := userbiz.NewRegisterBiz(userStore, bcrypt)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
