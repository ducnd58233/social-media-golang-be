package userbiz

import (
	"context"
	"social-media-be/common"
	"social-media-be/components/tokenprovider/hasher"
	usermodel "social-media-be/modules/user/model"
)

type RegiserStore interface {
	Find(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type registerBiz struct {
	registerStore RegiserStore
	hasher        hasher.Hasher
}

func NewRegisterBiz(registerStore RegiserStore, hasher hasher.Hasher) *registerBiz {
	return &registerBiz{
		registerStore: registerStore,
		hasher:        hasher,
	}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, err := biz.registerStore.Find(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	if err != nil && err == common.ErrRecordNotFound {
		data.Password = biz.hasher.Hash(data.Password)
		data.Role = "user"
		data.Status = 1

		if err := biz.registerStore.Create(ctx, data); err != nil {
			return common.ErrCannotCreateEntity(usermodel.EntityName, err)
		}

		return nil
	}

	return common.ErrDB(err)
}
