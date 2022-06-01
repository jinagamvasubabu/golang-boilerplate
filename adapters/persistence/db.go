package adapters

import (
	"fmt"

	"github.com/jinagamvasubabu/JITScheduler/config"
	"github.com/jinagamvasubabu/JITScheduler/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	cfg := config.GetConfig()
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.PgUser, cfg.PgPassword, cfg.PgHost, cfg.PgPort, cfg.DB)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//Auto Migration
	if cfg.Migrate {
		db.AutoMigrate(&models.Book{})
	}

	return db
}
