package helpers

import (
	"github.com/jinzhu/gorm"
	// sqlite driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ExecGORM executes a callback with a gorm DB
func ExecGORM(callback func(_db *gorm.DB)) {
	dbReference := EnvOrDefault("DATABASE_URI", "test.db")
	var db, err = gorm.Open("sqlite3", dbReference)
	if err != nil {
		CustomPanic(err)
	}
	defer db.Close()

	callback(db)
}
