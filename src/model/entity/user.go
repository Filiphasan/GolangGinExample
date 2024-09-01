package entity

import (
	"go-gin-example/src/constant"
	"time"
)

type User struct {
	Id              int              `gorm:"primaryKey autoIncrement column:id not null type:int" json:"id"`
	Name            string           `gorm:"column:name not null type:varchar(120)" json:"name"`
	Surname         string           `gorm:"column:surname not null type:varchar(120)" json:"surname"`
	Email           string           `gorm:"column:email not null type:varchar(240)" json:"email"`
	Password        string           `gorm:"column:password not null type:varchar(400)" json:"password"`
	CreateAt        time.Time        `gorm:"column:createAt not null type:timestamp" json:"createAt"`
	UserPermissions []UserPermission `gorm:"foreignKey:UserId" json:"userPermissions"`
}

func (User) TableName() string {
	return constant.UserTableName
}
