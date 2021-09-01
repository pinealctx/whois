package db

import (
	"github.com/pinealctx/whois/api/model"
	"github.com/pinealctx/whois/pkg/store"
	"gorm.io/gorm"
	"time"
)

type GormWeStore struct {
	Base
	db *gorm.DB
}

func NewWeGormStore(db *gorm.DB) store.WeStore {
	return GormWeStore{db: db}
}

func (g GormWeStore) LoadOpenInfo(wk *model.WeChatKey) (*model.UWechat, error) {
	return model.GetByWechatOpen(g.db, wk.OpenID, wk.AppID, false)
}

func (g GormWeStore) LoadMiniAppInfo(wk *model.WeChatKey) (*model.UWechat, error) {
	return model.GetByWechatMiniApp(g.db, wk.OpenID, wk.AppID, false)
}

func (g GormWeStore) UpdateOpenInfo(wf *model.WechatInfo, now time.Time) error {
	return model.UpdateUWechatOpenInfo(g.db, wf, now)
}

func (g GormWeStore) UpdateMiniAppInfo(wf *model.WechatInfo, now time.Time) error {
	return model.UpdateUWechatMiniAppInfo(g.db, wf, now)
}

func (g GormWeStore) AddOpenInfo(wf *model.WechatInfo, now time.Time) error {
	return model.AddUWechatOpenNoUID(g.db, wf, now)
}

func (g GormWeStore) AddMiniAppInfo(wf *model.WechatInfo, now time.Time) error {
	return model.AddUWechatMiniAppNoUID(g.db, wf, now)
}

func (g GormWeStore) BindOpenUserID(wk *model.WeChatKey, userID int32, now time.Time) error {
	return model.UpdateUWechatOpenUID(g.db, wk.OpenID, wk.AppID, userID, now)
}

func (g GormWeStore) BindMiniAppUserID(wk *model.WeChatKey, userID int32, now time.Time) error {
	return model.UpdateUWechatMiniAppUID(g.db, wk.OpenID, wk.AppID, userID, now)
}
