package store

import (
	"github.com/yxtiblya/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// create db conntection and return *gorm.DB, error
func NewDB(database_url string) (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(postgres.Open(database_url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB.AutoMigrate(
		&models.Contact{},
		&models.Mailing{},
		&models.Message{},
	)

	return DB, nil
}
