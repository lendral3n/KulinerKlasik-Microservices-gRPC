package database

import (
	"authservice/app/config"
	"authservice/internal/repository"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, 
		cfg.DB_PASSWORD, 
		cfg.DB_HOSTNAME, 
		cfg.DB_PORT, 
		cfg.DB_NAME)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&repository.User{},
		&repository.Role{},
	)

	return DB
}