package database

import (
	"github.com/xtt28/neptune/database/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectSQLite3(filename string) (conn *gorm.DB, err error) {
	conn, err = gorm.Open(sqlite.Open(filename), &gorm.Config{TranslateError: true})
	conn.AutoMigrate(&model.UserProfile{})
	conn.AutoMigrate(&model.Permission{})
	conn.AutoMigrate(&model.Balance{})
	conn.AutoMigrate(&model.PvPStat{})
	conn.AutoMigrate(&model.Punishment{})
	return
}
