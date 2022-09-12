package models

import (
	"path"

	"suxin2017.com/server/constants"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func GetDatabasePath() string {
	sqlitePath := path.Join(constants.GetServerDir(), "db")
	constants.InitNotExistDir(sqlitePath)
	return path.Join(sqlitePath, "test.db")
}

func ConnectDatabase() {
	var err error
	databasePath := GetDatabasePath()
	DB, err = gorm.Open(sqlite.Open(databasePath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		panic("Failed to connect to database!")
	}
	DB.AutoMigrate(&Domain{})
	DB.AutoMigrate(&User{})

}
