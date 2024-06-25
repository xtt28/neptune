package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSQLite3(filename string) (conn *gorm.DB, err error) {
	conn, err = gorm.Open(sqlite.Open(filename), &gorm.Config{TranslateError: true})
	return
}