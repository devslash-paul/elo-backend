package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	// I don't know'
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ExportDB struct {
	*gorm.DB
}

var currentWrapper *ExportDB

func Connect() (*ExportDB, error) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=postgres sslmode=disable password=postgres")
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the database successfully")

	wrapper := &ExportDB{db}
	currentWrapper = wrapper
	return wrapper, nil
}

func (wrapper *ExportDB) Migrate(logmode bool) error {
	wrapper.LogMode(logmode)
	wrapper.AutoMigrate(&League{})
	wrapper.AutoMigrate(&LeagueConfig{})
	wrapper.AutoMigrate(&Event{})
	wrapper.AutoMigrate(&User{})
	wrapper.AutoMigrate(&Worker{})
	wrapper.AutoMigrate(&Game{})
	wrapper.AutoMigrate(&GameAdditionals{})

	// wrapper.AutoMigrate(&EloWorker{})
	return nil
}

func (wrapper *ExportDB) CloseDB() {
	fmt.Println("Closing the database")
	wrapper.Close()
}

type ELO struct {
	ID        uint
	DeletedAt *time.Time `xml:"-" json:"-"`
	CreatedAt *time.Time `xml:"-" json:"-"`
	UpdatedAt *time.Time `xml:"-" json:"-"`
	score     uint
}
