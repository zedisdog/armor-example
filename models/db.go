package models

import (
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(
		&Account{},
		&Role{},
		&Permission{},
	)
}
