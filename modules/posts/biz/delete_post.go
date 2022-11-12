package postbiz

import (
	"context"
	"errors"
	postmodel "social-media-be/modules/posts/model"
)

type DeletePostStore interface {
	Find(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*postmodel.Post, error)
	SoftDelete(
		ctx context.Context,
		id int,
	) error
}

type deletePostBiz struct {
	store DeletePostStore
}

func NewDeletePostBiz(store DeletePostStore) *deletePostBiz {
	return &deletePostBiz{store: store}
}

func (biz *deletePostBiz) DeletePostById(ctx context.Context, id int) error {
	data, err := biz.store.Find(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if data.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.SoftDelete(ctx, id); err != nil {
		return err
	}

	return nil
}
