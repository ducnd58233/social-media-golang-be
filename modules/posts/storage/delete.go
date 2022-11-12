package poststorage

import (
	"context"
	"social-media-be/common"
	postmodel "social-media-be/modules/posts/model"
)

func (s *sqlStore) SoftDelete(
	ctx context.Context,
	id int,
) error {
	db := s.db.Table(postmodel.Post{}.TableName())

	if err := db.Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
