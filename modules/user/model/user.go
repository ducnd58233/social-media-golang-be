package usermodel

import (
	"errors"
	"social-media-be/common"
	"social-media-be/components/tokenprovider"
	"time"
)

const EntityName = "User"
const TableName = "users"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string     `json:"email" gorm:"column:email;"`
	Password        string     `json:"-" gorm:"column:password;"`
	LastName        string     `json:"last_name" gorm:"column:last_name;"`
	FirstName       string     `json:"first_name" gorm:"column:first_name;"`
	Gender          string     `json:"gender" gorm:"column:gender;"`
	Phone           string     `json:"phone" gorm:"column:phone;"`
	Role            string     `json:"-" gorm:"column:role;"`
	LastSeen        *time.Time `json:"last_seen,omitempty" gorm:"last_seen;"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (u *User) GetLastSeen() *time.Time {
	return u.LastSeen
}

func (User) TableName() string {
	return TableName
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string `json:"email" gorm:"column:email;"`
	Password        string `json:"-" gorm:"column:password;"`
	LastName        string `json:"last_name" gorm:"column:last_name;"`
	FirstName       string `json:"first_name" gorm:"column:first_name;"`
	Role            string `json:"-" gorm:"column:role;"`
}

func (UserCreate) TableName() string {
	return "users"
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserLogin struct {
	Email    string     `json:"email" form:"email" gorm:"column:email;"`
	Password string     `json:"password" form:"password" gorm:"column:password;"`
	LastSeen *time.Time `json:"last_seen,omitempty" gorm:"last_seen;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrUsernameOrPasswordInvalid = common.NewUnauthorized(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewUnauthorized(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
)
