package store

import (
	"github.com/pinealctx/whois/api/model"
	"time"
)

type MobileStore interface {
	LoadMobile(phone *model.MobileInfo) (*model.UMobile, error)
	AddMobile(phone *model.MobileInfo, userID int32, now time.Time) error
	UpsertMobileUID(phone *model.MobileInfo, userID int32, now time.Time) error

	IsDupErr(err error) bool
	IsNotFoundErr(err error) bool
}
