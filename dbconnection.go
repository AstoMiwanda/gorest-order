package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDBConnection(dbname string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(dbname), &gorm.Config{})

	return
}
