package store

import (
	"github.com/pinealctx/whois/api/model"
	"time"
)

type WeStore interface {
	LoadOpenInfo(wk *model.WeChatKey) (*model.UWechat, error)
	LoadMiniAppInfo(wk *model.WeChatKey) (*model.UWechat, error)

	UpdateOpenInfo(wf *model.WechatInfo, now time.Time) error
	UpdateMiniAppInfo(wf *model.WechatInfo, now time.Time) error

	AddOpenInfo(wf *model.WechatInfo, now time.Time) error
	AddMiniAppInfo(wf *model.WechatInfo, now time.Time) error

	BindOpenUserID(wk *model.WeChatKey, uid int32, now time.Time) error
	BindMiniAppUserID(wk *model.WeChatKey, uid int32, now time.Time) error

	IsNotFoundErr(err error) bool
}
