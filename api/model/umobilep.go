package model

import (
	"bytes"
	"github.com/pinealctx/neptune/store/gormx"
	"gorm.io/gorm"
	"time"
)

const (
	//upsert we chat info(without user id) sql part
	mobileFields = " (mobile, area, user_id, state, created_at, updated_at)"
	mobileVals   = " (?, ?, ?, 0, ?, ?)"
	mobileSet    = "on duplicate key update user_id=?, updated_at=?"
)

//AddUMobile : add user mobile
func AddUMobile(db *gorm.DB, area, mobile string, userID int32, now time.Time) error {
	var um = &UMobile{
		Mobile:    mobile,
		Area:      area,
		UserID:    userID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return db.Create(um).Error
}

//UpsertUMobile : upsert user mobile
func UpsertUMobile(db *gorm.DB, area, mobile string, uid int32, now time.Time) error {
	var buf = bytes.NewBuffer(make([]byte, 512))
	_, _ = buf.WriteString("insert into ")
	_, _ = buf.WriteString(uMobileTableName)
	_, _ = buf.WriteString(mobileFields)
	_, _ = buf.WriteString(mobileVals)
	_, _ = buf.WriteString(mobileSet)
	return db.Exec(buf.String(), mobile, area, uid, now, now, uid, now).Error
}

//GetByMobile : get user mobile info by mobile
func GetByMobile(db *gorm.DB, area, mobile string, forUpdate bool) (*UMobile, error) {
	var u UMobile
	var err = gormx.ForUpdate(db, forUpdate).Table(uMobileTableName).
		Where("`mobile` = ? AND `area` = ?", mobile, area).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

//GetMobilesByUID : get user mobile list by user id
func GetMobilesByUID(db *gorm.DB, userID int32) ([]UMobile, error) {
	var us []UMobile
	var err = db.Table(uMobileTableName).
		Where("`user_id` = ?", userID).Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}
