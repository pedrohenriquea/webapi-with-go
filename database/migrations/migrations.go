package migrations

import (
	"github.com/pedrohenriquea/webapi-with-go/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Book{})
}
