package model

import "github.com/jinzhu/gorm"

var db *gorm.DB

// SetDatabase function in model to keep main.go slim and both easy to test
func SetDatabase(database *gorm.DB) {
	db = database
}

// User struct for DB connection test
type User struct {
	FirstName string
	LastName  string
}
