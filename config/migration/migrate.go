package migration

import (
	"github.com/jinzhu/gorm"
	"github.com/keiya01/ginsampleapp/src/models"
)

func DBmigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})
	db.AutoMigrate(&models.User{})
}
