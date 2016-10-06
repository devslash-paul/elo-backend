package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

var constDB *gorm.DB

func Connect() error {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=postgres sslmode=disable password=postgres")
	if err != nil {
		return err
	}

	constDB = db
	return nil
}

func Migrate() error {
	constDB.LogMode(true)
	constDB.AutoMigrate(&Product{})
	return nil
}

func Close() {
	defer constDB.Close()
}
