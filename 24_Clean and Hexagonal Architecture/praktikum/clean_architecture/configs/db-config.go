package config

import (
	model "clean_hexagonal_architecture/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("app.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = MigrateDB(db)
	if err != nil {
		panic(err)
	}

	return db, err
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}
