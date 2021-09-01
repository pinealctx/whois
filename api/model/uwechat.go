package model

import (
	"time"
)

var (
	uWechatOpenTableName    = "user_wechat_mini_app"
	uWechatMiniAppTableName = "user_wechat_open"
)

type WeChatKey struct {
	OpenID string
	AppID  string
}

func (m *WeChatKey) Key() string {
	return PinWechatKey(m.OpenID, m.AppID)
}

type WechatInfo struct {
	OpenID string
	AppID  string

	NickName   string
	Country    string
	Province   string
	City       string
	HeadImgUrl string

	Sex int32
}

func (m *WechatInfo) Key() string {
	return PinWechatKey(m.OpenID, m.AppID)
}

func (m *WechatInfo) WeKey() *WeChatKey {
	return &WeChatKey{
		OpenID: m.OpenID,
		AppID:  m.AppID,
	}
}

type UWechat struct {
	OpenID string `gorm:"column:open_id"`
	AppID  string `gorm:"column:app_id"`

	UserID int32 `gorm:"column:user_id"`
	Sex    int32 `gorm:"column:sex"`
	State  int32 `gorm:"column:state"`

	NickName   string `gorm:"column:nick_name"`
	Country    string `gorm:"column:country"`
	Province   string `gorm:"column:province"`
	City       string `gorm:"column:city"`
	HeadImgUrl string `gorm:"column:head_img_url"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *UWechat) Size() int {
	return 1
}

func (u *UWechat) HasNoChange(m *WechatInfo) bool {
	return u.OpenID == m.OpenID &&
		u.AppID == m.AppID &&
		u.NickName == m.NickName &&
		u.Country == m.Country &&
		u.Province == m.Province &&
		u.City == m.City &&
		u.HeadImgUrl == m.HeadImgUrl &&
		u.Sex == m.Sex
}

func (u *UWechat) UpdFrom(w *WechatInfo, now time.Time) {
	u.OpenID = w.OpenID
	u.AppID = w.AppID
	u.NickName = w.NickName
	u.Country = w.Country
	u.Province = w.Province
	u.City = w.City
	u.HeadImgUrl = w.HeadImgUrl
	u.Sex = w.Sex
	u.UpdatedAt = now
}

func SetUWechatMiniAppTableName(name string) {
	uWechatMiniAppTableName = name
}

func SetUWechatOpenTableName(name string) {
	uWechatOpenTableName = name
}
