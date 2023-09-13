package database

import (
	"fmt"
	"time"
	"zsocket_service/config"
	"zsocket_service/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBInstance struct {
	DB *gorm.DB
}

func Connection(configdb config.Credentials) (*DBInstance, error) {
	var (
		db    *gorm.DB
		err   error
		count int = 0
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		configdb.User,
		configdb.Password,
		configdb.Host,
		configdb.Port,
		configdb.Database)

	for {
		if count == 10 {
			break
		}
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			count++
			time.Sleep(5 * time.Second)
			continue
		} else {
			return &DBInstance{DB: db}, err
		}
	}

	return &DBInstance{DB: db}, err
}

func Migrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.TopicMaster{},
		&model.SessionMaster{})
}
