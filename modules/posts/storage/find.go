package poststorage

import (
	"context"
	"social-media-be/common"
	postmodel "social-media-be/modules/posts/model"

	"gorm.io/gorm"
)

func (s *sqlStore) Find(
	ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string,
) (*postmodel.Post, error) {
	var result postmodel.Post

	if err := s.db.Where(conditions).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
