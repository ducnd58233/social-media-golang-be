package uploadstorage

import (
	"context"
	"social-media-be/common"

	"gorm.io/gorm"
)

func (s *sqlStore) Find(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*common.Image, error) {
	db := s.db.Table(common.Image{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var image common.Image

	if err := db.Where(conditions).First(&image).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &image, nil
}
