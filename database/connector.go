package database

import (
	"github.com/xtt28/neptune/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSQLite3(filename string) (conn *gorm.DB, err error) {
	conn, err = gorm.Open(sqlite.Open(filename), &gorm.Config{TranslateError: true})
	conn.AutoMigrate(&models.Permission{})
	return
}
