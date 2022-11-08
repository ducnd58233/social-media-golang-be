package userbiz

import (
	"context"
	"social-media-be/common"
	"social-media-be/components/tokenprovider"
	"social-media-be/components/tokenprovider/hasher"
	usermodel "social-media-be/modules/user/model"
	"time"
)

type LoginStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	UpdateTime(ctx context.Context, time time.Time, email string) error
}

type loginBiz struct {
	store         LoginStore
	tokenProvider tokenprovider.Provider
	hasher        hasher.Hasher
	expiry        int
}

func NewLoginBiz(store LoginStore, tokenProvider tokenprovider.Provider, hasher hasher.Hasher, expiry int) *loginBiz {
	return &loginBiz{
		store:         store,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Account, error) {
	user, err := biz.store.Find(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	match := biz.hasher.Validate(data.Password, user.Password)

	if !match {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	biz.store.UpdateTime(ctx, time.Now().UTC(), user.Email)

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	account := usermodel.NewAccount(accessToken, refreshToken)

	return account, nil
}
