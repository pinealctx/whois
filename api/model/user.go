package model

import (
	"time"
)

var (
	userTableName = "user"
)

type User struct {
	ID        int32     `gorm:"column:id"`
	State     int32     `gorm:"column:state"`
	NickName  string    `gorm:"column:nick_name"`
	Avatar    string    `gorm:"column:avatar"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u User) TableName() string {
	return userTableName
}

func (u *User) Size() int {
	return 1
}

func SetUserTableName(name string) {
	userTableName = name
}
