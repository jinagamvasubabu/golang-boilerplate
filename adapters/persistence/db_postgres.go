package persistence

import (
	"fmt"
	"time"

	Logger "github.com/jinagamvasubabu/golang-boilerplate/adapters/logger"
	"github.com/jinagamvasubabu/golang-boilerplate/config"
	"github.com/jinagamvasubabu/golang-boilerplate/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitPostgresDatabase() (*gorm.DB, error) {
	cfg := config.GetConfig()
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.PgUser, cfg.PgPassword, cfg.PgHost, cfg.PgPort, cfg.DB)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		Logger.Errorf("err:=%s", err.Error())
		time.Sleep(10 * time.Second)
		InitPostgresDatabase()
	}

	//Auto Migration
	if cfg.Migrate {
		db.AutoMigrate(&model.Book{})
	}
	// Get generic database object sql.DB to use its functions
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(cfg.MaxOpenConnections - 1)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnections)
	Logger.Info("Successfully created db connection")
	return db, nil
}

func GetDBConn() *gorm.DB {
	return db
}

func SetDBConn(conn *gorm.DB) {
	db = conn
}
