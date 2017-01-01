package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var currentWrapper DB

type DB interface {
	Migrate(bool) error
	CloseDB()

	// Database functions

	PreloadAllFor(interface{}, string)
	Create(interface{})
	Delete(interface{}, ...interface{})
	First(interface{}, ...interface{})
	Where(interface{}, string, ...interface{})
	PrimaryWithCondidtion(interface{}, uint64, string, ...interface{})
	GetRelatedFor(interface{}, ...interface{})
}

type postgresDB struct {
	db *gorm.DB
}

func Connect() (DB, error) {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=postgres sslmode=disable password=postgres")
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the database successfully")

	wrapper := &postgresDB{db}
	currentWrapper = wrapper
	return wrapper, nil
}
func (s *postgresDB) Delete(del interface{}, where ...interface{}) {
	s.db.Delete(del, where)
}

func (s *postgresDB) GetRelatedFor(loaded interface{}, related ...interface{}) {
	model := s.db.Model(loaded)
	for _, rel := range related {
		model.Related(rel)
	}
}

func (s *postgresDB) Where(out interface{}, query string, values ...interface{}) {
	s.db.Where(query, values).Find(out)
}

func (s *postgresDB) PrimaryWithCondidtion(out interface{}, id uint64, query string, values ...interface{}) {
	s.db.Where(query, values).Find(out, id)
}
func (s *postgresDB) PreloadAllFor(value interface{}, col string) {
	s.db.Preload(col).Find(value)
}

func (s *postgresDB) Create(value interface{}) {
	s.db.Create(value)
}

func (s *postgresDB) First(out interface{}, where ...interface{}) {
	s.db.First(out, where)
}

func (wrapper *postgresDB) Migrate(logmode bool) error {
	wrapper.db.LogMode(logmode).
		AutoMigrate(&League{}).
		AutoMigrate(&LeagueConfig{}).
		AutoMigrate(&Event{}).
		AutoMigrate(&User{}).
		AutoMigrate(&Worker{}).
		AutoMigrate(&Game{}).
		AutoMigrate(&GameAdditionals{})

	// wrapper.AutoMigrate(&EloWorker{})
	return nil
}

func (wrapper *postgresDB) CloseDB() {
	fmt.Println("Closing the database")
	wrapper.db.Close()
}

type ELO struct {
	ID        uint
	DeletedAt *time.Time `xml:"-" json:"-"`
	CreatedAt *time.Time `xml:"-" json:"-"`
	UpdatedAt *time.Time `xml:"-" json:"-"`
	score     uint
}
