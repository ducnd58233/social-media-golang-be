package common

import "time"

const (
	DbTypeUser = iota
	DbTypePost
	DbTypeComment
	DbTypeReaction
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
	GetLastSeen() *time.Time
}