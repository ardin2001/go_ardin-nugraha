package utils

import (
	"fmt"
	"orm_mvc/config"
	"orm_mvc/models"
)

func UserMigrate() {
	DB, err := config.InitDB()
	if err != nil {
		fmt.Println("Failed connect to database : ", err.Error())
		return
	}
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
