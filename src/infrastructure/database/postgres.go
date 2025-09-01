package database

import (
	"log"
	"time"

	"github.com/faujiahmat/zentra-user-service/src/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres() *gorm.DB {
	dsn := config.Conf.Postgres.Dsn

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)

	return db
}

func ClosePostgres(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
