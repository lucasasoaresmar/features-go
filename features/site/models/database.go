package models

import (
	"github.com/jinzhu/gorm"
	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db, _ = gorm.Open("sqlite3", "test.db")

func execute(callback func(_db *gorm.DB)) {
	callback(db)
}
