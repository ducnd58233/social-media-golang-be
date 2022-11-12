package userstorage

import (
	"context"
	"social-media-be/common"
	usermodel "social-media-be/modules/user/model"
	"time"
)

func (s *sqlStore) UpdateTime(
	ctx context.Context, 
	time time.Time, 
	email string,
) error {
	db := s.db.Begin().Table(usermodel.TableName)

	if err := db.Where("email = ?", email).Update("last_seen", time).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
