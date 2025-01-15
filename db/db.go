package db

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/juliofilizzola/bot_discord/application/domain/model"
	"github.com/juliofilizzola/bot_discord/config/env"
	_ "github.com/lib/pq"
)


func ConnectDB() (*gorm.DB, error) {

	var err error
	var db *gorm.DB
	db, err = gorm.Open(env.DbType, env.DatabaseURL)

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if env.AutoMigrateDb == "true" {
		db.AutoMigrate(&model.PR{}, &model.User{})
	}
	return db, nil

}
