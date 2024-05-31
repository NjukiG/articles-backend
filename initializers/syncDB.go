package initializers

import (
	"articles-api/models"
)

func SyncDb() {
	// Sync the postgrs DB to create tables for the models
	DB.AutoMigrate(&models.User{}, &models.Article{}, &models.Comment{})
}
