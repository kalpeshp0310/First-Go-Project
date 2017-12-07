package models

import (
	"github.com/jinzhu/gorm"
	// needed for database drivers
	_ "github.com/lib/pq"
)

var db *gorm.DB

// ConnectDB returns a gorm DB connection
func ConnectDB() *gorm.DB {
	dbase, err := gorm.Open("postgres", "postgres:///hellogo_development?sslmode=disable")
	if err != nil {
		panic(err)
	}
	db = dbase
	return dbase
}
