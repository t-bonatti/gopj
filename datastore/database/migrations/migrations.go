package migrations

import (
	"t-bonatti/gopj/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Company{})
}
