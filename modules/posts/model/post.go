package postmodel

import "social-media-be/common"

const EntityName = "Post"
const TableName = "posts"

type Post struct {
	common.SQLModel `json:",inline"`
	FakeOwnerId     *common.UID           `json:"owner_id" gorm:"-"`
	OwnerId         int            `json:"-" gorm:"column:owner_id"`
	Images          *common.Images `json:"images" gorm:"column:images"`
	Desc            string         `json:"desc" gorm:"column:desc"`
}

func (Post) TableName() string {
	return TableName
}

func (p *Post) Mask(isAdmin bool) {
	p.GenUID(common.DbTypePost)
}

func (p *Post) MaskOwnerId(isAdmin bool) {
	uid := common.NewUID(uint32(p.OwnerId), common.DbTypeUser, 1)
	p.FakeOwnerId = &uid
}

type PostNoImages struct {
	common.SQLModel `json:",inline"`
	OwnerId         int    `json:"owner_id" gorm:"column:owner_id"`
	Desc            string `json:"desc" gorm:"column:desc"`
}

func (PostNoImages) TableName() string {
	return Post{}.TableName()
}
