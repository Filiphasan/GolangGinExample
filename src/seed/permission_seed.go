package seed

import (
	"go-gin-example/src/model/entity"
	"gorm.io/gorm"
	"time"
)

func CreateBasePermissions(db *gorm.DB) {
	// Check permission table any data
	var count int64
	db.Model(&entity.Permission{}).Count(&count)
	if count > 0 {
		return
	}
	// Create permission
	db.Create(&entity.Permission{
		Name:     "regular",
		CreateAt: time.Now(),
	})
}
