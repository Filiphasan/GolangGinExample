package entity

import (
	"go-gin-example/src/constant"
	"time"
)

type Permission struct {
	Id              int              `gorm:"primaryKey autoIncrement column:id not null type:int" json:"id"`
	Name            string           `gorm:"column:name not null type:varchar(120)" json:"name"`
	CreateAt        time.Time        `gorm:"column:createAt not null type:timestamp" json:"createAt"`
	UserPermissions []UserPermission `gorm:"foreignKey:PermissionId" json:"userPermissions"`
}

func (Permission) TableName() string {
	return constant.PermissionTableName
}
