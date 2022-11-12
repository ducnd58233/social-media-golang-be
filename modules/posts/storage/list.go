package poststorage

import (
	"context"
	"social-media-be/common"
	postmodel "social-media-be/modules/posts/model"
)

func (s *sqlStore) List(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) ([]postmodel.Post, error) {
	var result []postmodel.Post

	db := s.db.Table(postmodel.Post{}.TableName()).Where(conditions)

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if err := db.
		Order("created_at desc").
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
