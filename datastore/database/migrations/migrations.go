package migrations

import (
	"t-bonatti/gopj/model"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.Company{})
}
