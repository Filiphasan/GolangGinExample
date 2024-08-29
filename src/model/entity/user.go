package entity

import (
	"go-gin-example/src/constant"
	"time"
)

type User struct {
	Id       int       `gorm:"primaryKey autoIncrement column:id not null" json:"id"`
	Name     string    `gorm:"column:name not null" json:"name"`
	Email    string    `gorm:"column:email not null" json:"email"`
	Password string    `gorm:"column:password not null" json:"password"`
	CreateAt time.Time `gorm:"column:createAt not null" json:"createAt"`
}

func (User) TableName() string {
	return constant.UserTableName
}
