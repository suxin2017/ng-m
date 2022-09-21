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
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to database!")
	}
	DB.AutoMigrate(&Domain{})
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Location{})
	DB.AutoMigrate(&Server{})
	DB.AutoMigrate(&Upstream{})

	if constants.IsDev() {

		admin := &User{
			Email: "abc@123.com", Password: "123123",
			Domain: []Domain{
				{Domain: "www.test.com", Desc: "测试域名", CreatedUserID: 1, UpdatedUserId: 1},
			},
		}
		DB.First(admin, 1)
		if admin.ID == 0 {

			DB.Create(admin)
		}
		// DB.Model()
		// DB.FirstOrCreate(&Domain{crea})
	}
}
