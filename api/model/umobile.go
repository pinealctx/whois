package model

import (
	"time"
)

var (
	uMobileTableName = "user_mobile"
)

type UMobile struct {
	Mobile    string    `gorm:"column:mobile"`
	Area      string    `gorm:"column:area"`
	UserID    int32     `gorm:"column:user_id"`
	State     int32     `gorm:"column:state"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u UMobile) TableName() string {
	return uMobileTableName
}

func (u *UMobile) Size() int {
	return 1
}

type MobileInfo struct {
	Area   string
	Mobile string
}

func (m *MobileInfo) Key() string {
	return PinMobile(m.Mobile, m.Area)
}

func SetUMobileTableName(name string) {
	uMobileTableName = name
}
