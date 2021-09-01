package store

import (
	"github.com/pinealctx/whois/api/model"
	"time"
)

type UserStore interface {
	LoadUser(uid int32) (*model.User, error)
	LoadUserMobiles(uid int32) ([]model.MobileInfo, error)

	AddUser(uid int32, nickName, avatar string, now time.Time) error

	UpdateNick(uid int32, nickName string, now time.Time) error
	UpdateAvatar(uid int32, avatar string, now time.Time) error

	IsDupErr(err error) bool
}
