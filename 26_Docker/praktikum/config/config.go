package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() (*gorm.DB, error) {

	config := Config{
		DB_Username: "root",
		DB_Password: "123456",
		DB_Port:     "3307",
		DB_Host:     "host.docker.internal",
		DB_Name:     "gorm_crud",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	DB, err := gorm.Open("mysql", connectionString)
	// DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return DB, nil
}
