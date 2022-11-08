package userstorage

import (
	"context"
	"social-media-be/common"
	usermodel "social-media-be/modules/user/model"
)

func (s *sqlStore) Create(ctx context.Context, data *usermodel.UserCreate) error {
	db := s.db.Begin().Table(data.TableName())

	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
