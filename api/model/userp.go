package model

import (
	"github.com/pinealctx/neptune/store/gormx"
	"gorm.io/gorm"
	"time"
)

//AddUser : add user
func AddUser(db *gorm.DB, id int32, nickName string, avatar string, now time.Time) error {
	var u = &User{
		ID:        id,
		NickName:  nickName,
		Avatar:    avatar,
		CreatedAt: now,
		UpdatedAt: now,
	}
	return db.Create(u).Error
}

//UpdateNickName : update user nickname
func UpdateNickName(db *gorm.DB, id int32, nickName string, now time.Time) error {
	return db.Table(userTableName).Where("`id` = ?", id).Updates(
		map[string]interface{}{
			"nick_name":  nickName,
			"updated_at": now,
		}).Error
}

//UpdateAvatar : update user avatar
func UpdateAvatar(db *gorm.DB, id int32, avatar string, now time.Time) error {
	return db.Table(userTableName).Where("`id` = ?", id).Updates(
		map[string]interface{}{
			"avatar":     avatar,
			"updated_at": now,
		}).Error
}

//GetUserByID : get user by id
func GetUserByID(db *gorm.DB, id int32, forUpdate bool) (*User, error) {
	var u User
	var err = gormx.ForUpdate(db, forUpdate).Table(userTableName).Where("`id` = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

//GetLastUID : get last user id
func GetLastUID(db *gorm.DB) (int32, error) {
	var us []User
	var err = db.Table(userTableName).Order("id desc").Limit(1).Find(&us).Error
	if err != nil {
		return 0, err
	}
	if len(us) == 0 {
		return 0, nil
	}
	return us[0].ID, nil
}
