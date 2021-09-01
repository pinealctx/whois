package model

import (
	"github.com/pinealctx/neptune/store/gormx"
	"gorm.io/gorm"
	"time"
)

/*******************WeChat Open DB functions********************/

//AddUWechatOpenNoUID : add we chat open info without user id
func AddUWechatOpenNoUID(db *gorm.DB, w *WechatInfo, now time.Time) error {
	return addUWechatInfoNoUID(db, uWechatOpenTableName, w, now)
}

//UpdateUWechatOpenInfo : update we chat open info
func UpdateUWechatOpenInfo(db *gorm.DB, w *WechatInfo, now time.Time) error {
	return updateUWechatInfo(db, uWechatOpenTableName, w, now)
}

//UpdateUWechatOpenUID : update user id of we chat open info
func UpdateUWechatOpenUID(db *gorm.DB, openID, appID string, uid int32, now time.Time) error {
	return updateUWechatUID(db, uWechatOpenTableName, openID, appID, uid, now)
}

//GetByWechatOpen : get we chat open info
func GetByWechatOpen(db *gorm.DB, openID, appID string, forUpdate bool) (*UWechat, error) {
	return getByWechat(db, uWechatOpenTableName, openID, appID, forUpdate)
}

/*******************WeChat Mini App DB functions********************/

//AddUWechatMiniAppNoUID : add we chat mini app info without user id
func AddUWechatMiniAppNoUID(db *gorm.DB, w *WechatInfo, now time.Time) error {
	return addUWechatInfoNoUID(db, uWechatMiniAppTableName, w, now)
}

//UpdateUWechatMiniAppInfo : update we chat mini app info
func UpdateUWechatMiniAppInfo(db *gorm.DB, w *WechatInfo, now time.Time) error {
	return updateUWechatInfo(db, uWechatMiniAppTableName, w, now)
}

//UpdateUWechatMiniAppUID : update user id of we chat mini app info
func UpdateUWechatMiniAppUID(db *gorm.DB, openID, appID string, uid int32, now time.Time) error {
	return updateUWechatUID(db, uWechatMiniAppTableName, openID, appID, uid, now)
}

//GetByWechatMiniApp : get we chat mini app info
func GetByWechatMiniApp(db *gorm.DB, openID, appID string, forUpdate bool) (*UWechat, error) {
	return getByWechat(db, uWechatMiniAppTableName, openID, appID, forUpdate)
}

//addWechatInfo : add we chat info without user id
func addUWechatInfoNoUID(db *gorm.DB, tableName string, w *WechatInfo, now time.Time) error {
	var uw = &UWechat{}
	uw.UpdFrom(w, now)
	uw.CreatedAt = now
	return db.Table(tableName).Select("open_id", "app_id",
		"sex",
		"nick_name", "country", "province", "city", "head_img_url",
		"created_at", "updated_at").Create(uw).Error
}

//updateUWechatInfo : update user we chat info
func updateUWechatInfo(db *gorm.DB, tableName string, w *WechatInfo, now time.Time) error {
	return db.Table(tableName).Where("`open_id` = ? AND `app_id` = ?", w.OpenID, w.AppID).Updates(
		map[string]interface{}{
			"sex":          w.Sex,
			"nick_name":    w.NickName,
			"country":      w.Country,
			"province":     w.Province,
			"city":         w.City,
			"head_img_url": w.HeadImgUrl,
			"updated_at":   now,
		}).Error
}

//updateWechatUID : update user we chat related user id
func updateUWechatUID(db *gorm.DB, tableName string, openID, appID string, uid int32, now time.Time) error {
	return db.Table(tableName).Where("`open_id` = ? AND `app_id` = ?", openID, appID).Updates(
		map[string]interface{}{
			"user_id":    uid,
			"updated_at": now,
		}).Error
}

//getByWechat : get user we chat info by we chat
func getByWechat(db *gorm.DB, tableName string, openID, appID string, forUpdate bool) (*UWechat, error) {
	var u UWechat
	var err = gormx.ForUpdate(db, forUpdate).Table(tableName).
		Where("`open_id` = ? AND `app_id` = ?", openID, appID).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}
