package db

import (
	"github.com/pinealctx/whois/api/model"
	"github.com/pinealctx/whois/pkg/store"
	"gorm.io/gorm"
	"time"
)

type GormUStore struct {
	Base
	db *gorm.DB
}

func NewMobUStore(db *gorm.DB) store.UserStore {
	return GormUStore{db: db}
}

func (g GormUStore) LoadUser(uid int32) (*model.User, error) {
	return model.GetUserByID(g.db, uid, false)
}

func (g GormUStore) LoadUserMobiles(uid int32) ([]model.MobileInfo, error) {
	var mobs, err = model.GetMobilesByUID(g.db, uid)
	if err != nil {
		return nil, err
	}
	var l = len(mobs)
	if l == 0 {
		return nil, nil
	}
	var rs = make([]model.MobileInfo, l)
	for i := 0; i < l; i++ {
		rs[i].Mobile, rs[i].Area = mobs[i].Mobile, mobs[i].Area
	}
	return rs, nil
}

func (g GormUStore) AddUser(uid int32, nickName, avatar string, now time.Time) error {
	return model.AddUser(g.db, uid, nickName, avatar, now)
}

func (g GormUStore) UpdateNick(uid int32, nickName string, now time.Time) error {
	return model.UpdateNickName(g.db, uid, nickName, now)
}

func (g GormUStore) UpdateAvatar(uid int32, avatar string, now time.Time) error {
	return model.UpdateAvatar(g.db, uid, avatar, now)
}
