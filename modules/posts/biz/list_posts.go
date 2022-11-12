package postbiz

import (
	"context"
	"social-media-be/common"
	postmodel "social-media-be/modules/posts/model"
)

type ListPostsStore interface {
	List(
		ctx context.Context,
		conditions map[string]interface{},
		// filter *postmodel.Filter,
		moreKeys ...string,
	) ([]postmodel.Post, error)
}

type listPostsBiz struct {
	store ListPostsStore
}

func NewListPostsBiz(store ListPostsStore) *listPostsBiz {
	return &listPostsBiz{store: store}
}

func (biz *listPostsBiz) ListPost(ctx context.Context) ([]postmodel.Post, error) {
	result, err := biz.store.List(ctx, map[string]interface{}{})

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	return result, nil
}

func (biz *listPostsBiz) ListPostByUserId(ctx context.Context, id int) ([]postmodel.Post, error) {
	result, err := biz.store.List(ctx, map[string]interface{}{"owner_id": id})

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	return result, nil
}