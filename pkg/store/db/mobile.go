package db

import (
	"github.com/pinealctx/whois/api/model"
	"github.com/pinealctx/whois/pkg/store"
	"gorm.io/gorm"
	"time"
)

type GormMobStore struct {
	Base
	db *gorm.DB
}

func NewMobGormStore(db *gorm.DB) store.MobileStore {
	return GormMobStore{db: db}
}

func (g GormMobStore) LoadMobile(m *model.MobileInfo) (*model.UMobile, error) {
	return model.GetByMobile(g.db, m.Area, m.Mobile, false)
}

func (g GormMobStore) AddMobile(m *model.MobileInfo, uid int32, now time.Time) error {
	return model.AddUMobile(g.db, m.Area, m.Mobile, uid, now)
}

func (g GormMobStore) UpsertMobileUID(m *model.MobileInfo, uid int32, now time.Time) error {
	return model.UpsertUMobile(g.db, m.Area, m.Mobile, uid, now)
}
