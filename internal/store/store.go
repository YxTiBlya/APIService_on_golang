package store

import (
	"fmt"

	"github.com/yxtiblya/internal/cfg"
	"github.com/yxtiblya/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// create db conntection and return *gorm.DB, error
func NewDB(c *cfg.Config) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=%v", c.Postgres_user, c.Postgres_password, c.Postgres_host, c.Postgres_port, c.Postgres_dbname, c.Postgres_ssl)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
