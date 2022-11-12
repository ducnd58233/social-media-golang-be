package postbiz

import (
	"context"
	"social-media-be/common"
	postmodel "social-media-be/modules/posts/model"
)

type GetPostStore interface {
	Find(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*postmodel.Post, error)
}

type getPostBiz struct {
	store GetPostStore
}

func NewGetPostBiz(store GetPostStore) *getPostBiz {
	return &getPostBiz{store: store}
}

func (biz *getPostBiz) GetPost(ctx context.Context, id int) (*postmodel.Post, error) {
	data, err := biz.store.Find(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.ErrRecordNotFound {
			return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	return data, nil
}
