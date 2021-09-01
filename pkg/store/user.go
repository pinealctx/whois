package store

import (
	"github.com/pinealctx/whois/api/model"
	"time"
)

type UserStore interface {
	LoadUser(userID int32) (*model.User, error)
	LoadUserMobiles(userID int32) ([]model.MobileInfo, error)

	AddUser(userID int32, nickName, avatar string, now time.Time) error

	UpdateNick(userID int32, nickName string, now time.Time) error
	UpdateAvatar(userID int32, avatar string, now time.Time) error

	IsDupErr(err error) bool
}
