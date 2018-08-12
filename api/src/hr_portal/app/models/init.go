package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Global database references
var db *gorm.DB

// Database settings
var db_name = "hr_portal"
var db_user = "root"
var db_pw = ""

// Create database connection
func Init_DB() {
	db, err := gorm.Open("mysql", db_user + ":" + db_pw + "@/" + db_name + "?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.Close()
}