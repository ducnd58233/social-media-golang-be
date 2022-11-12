package postbiz

import (
	"context"
	"errors"
	postmodel "social-media-be/modules/posts/model"
)

type UpdatePostStore interface {
	Find(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*postmodel.Post, error)
	Update(
		ctx context.Context,
		id int,
		data *postmodel.Post,
	) error
}

type updatePostBiz struct {
	store UpdatePostStore
}

func NewUpdatePostBiz(store UpdatePostStore) *updatePostBiz {
	return &updatePostBiz{store: store}
}

func (biz *updatePostBiz) UpdatePostById(ctx context.Context, id int, data *postmodel.Post) error {
	oldData, err := biz.store.Find(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
