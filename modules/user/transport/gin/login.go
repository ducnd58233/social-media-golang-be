package usergin

import (
	"net/http"
	"social-media-be/common"
	component "social-media-be/components"
	"social-media-be/components/tokenprovider/hasher"
	"social-media-be/components/tokenprovider/jwt"
	userbiz "social-media-be/modules/user/biz"
	usermodel "social-media-be/modules/user/model"
	userstorage "social-media-be/modules/user/storage"

	"github.com/gin-gonic/gin"
)

func LoginHandler(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		bcrypt := hasher.NewBcryptHash()

		biz := userbiz.NewLoginBiz(store, tokenProvider, bcrypt, 60*60*24*7)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
