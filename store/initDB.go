package store

import (
	"BookManagement/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var _db *gorm.DB

func Init(url string) *gorm.DB {
	if _db != nil {
		return _db
	}
	_db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	err = _db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil
	}

	return _db
}
