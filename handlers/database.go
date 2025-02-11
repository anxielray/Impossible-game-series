package handlers

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var Db *gorm.DB 

	
// User represents the users table structure

func CreateMoveTable(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}

func StartDbConn() error {
	var err error
	Db, err = gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := Db.DB()
	if err != nil {
		return err
	}

	// Check the connection
	if err = sqlDB.Ping(); err != nil {
		return err
	}

	log.Println("Database connected successfully")
	return nil
}
