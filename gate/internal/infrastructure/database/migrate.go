package database

import (
	"gorm.io/gorm"
)

func AutoMigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
	//&domain.User{},
	// 其他 model...
	)
}
