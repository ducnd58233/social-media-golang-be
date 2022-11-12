package postbiz

import (
	"context"
	postmodel "social-media-be/modules/posts/model"
)

type CreatePostStore interface {
	Create(ctx context.Context, data *postmodel.Post) error
}

type createPostBiz struct {
	store CreatePostStore
}

func NewCreatePostBiz(store CreatePostStore) *createPostBiz {
	return &createPostBiz{store: store}
}

func (biz *createPostBiz) CreatePost(ctx context.Context, data *postmodel.Post) error {
	err := biz.store.Create(ctx, data)

	return err
}