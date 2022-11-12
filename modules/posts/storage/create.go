package poststorage

import (
	"context"
	"social-media-be/common"
	postmodel "social-media-be/modules/posts/model"
)

func (s *sqlStore) Create(ctx context.Context, data *postmodel.Post) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
