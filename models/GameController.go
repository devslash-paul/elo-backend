package models

import (
	"errors"
	"time"
)

type GameController struct {
	db *ExportDB
}

type Game struct {
	ID              uint
	DeletedAt       *time.Time `xml:"-" json:"-"`
	CreatedAt       *time.Time `xml:"-" json:"-"`
	UpdatedAt       *time.Time `xml:"-" json:"-"`
	WinnerID        uint
	Winner          User `gorm:"ForeignKey:WinnerID"`
	LoserID         uint
	Loser           User              `gorm:"ForeignKey:LoserID"`
	League          League            `gorm:"ForeignKey:LeagueID"` // use LeagueID as foreign key
	LeagueID        uint              `json:"-" xml:"-"`
	GameAdditionals []GameAdditionals `json:",omitempty" xml:",omitempty"`
}

type GameAdditionals struct {
	ID        uint
	DeletedAt *time.Time `xml:"-" json:"-"`
	CreatedAt *time.Time `xml:"-" json:"-"`
	UpdatedAt *time.Time `xml:"-" json:"-"`
	key       string
	value     string
}

func NewGameController(db *ExportDB) *GameController {
	return &GameController{db}
}

func (gc *GameController) GetGameForEvent(event Event) (*Game, error) {
	game := Game{}
	gc.db.First(&game, event.RelatedEventID)

	if game.ID == 0 {
		return nil, errors.New("No game has been found with that ID")
	}

	return &game, nil
}

func (gc *GameController) Create(game *Game) {
	gc.db.Create(game)
}

func (gc *GameController) GetById(league *League, id uint64) (*Game, error) {
	game := &Game{}
	gc.db.Where("League = ?", league.ID).First(game, id)
	return game, nil
}
