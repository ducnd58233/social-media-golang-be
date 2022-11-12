package poststorage

import (
	"context"
	"social-media-be/common"
	postmodel "social-media-be/modules/posts/model"
)

func (s *sqlStore) Update(
	ctx context.Context,
	id int,
	data *postmodel.Post,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
