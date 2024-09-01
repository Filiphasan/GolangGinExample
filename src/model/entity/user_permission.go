package entity

import (
	"go-gin-example/src/constant"
	"time"
)

type UserPermission struct {
	UserId       int        `gorm:"column:userId not null type:int" json:"userId"`
	PermissionId int        `gorm:"column:permissionId not null type:int" json:"permissionId"`
	CreateAt     time.Time  `gorm:"column:createAt not null type:timestamp" json:"createAt"`
	User         User       `gorm:"foreignKey:UserId" json:"user"`
	Permission   Permission `gorm:"foreignKey:PermissionId" json:"permission"`
}

func (UserPermission) TableName() string {
	return constant.UserPermissionTableName
}
