package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type GameController struct {
	db DB
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

func (g *Game) AfterCreate(scope *gorm.Scope) error {
	ev := Event{
		EventName:      EVENT_GAME_PLAYED,
		RelatedTable:   "Game",
		RelatedEventID: g.ID,
	}
	scope.DB().Create(&ev)
	return nil
}

func NewGameController(db DB) *GameController {
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
	// ensure that a db instance is privately kept for use
	currentWrapper = gc.db
	gc.db.Create(game)
}

func (gc *GameController) GetById(league *League, id uint64) (*Game, error) {
	game := &Game{}
	gc.db.PrimaryWithCondidtion(game, id, "League = ?", league.ID)
	return game, nil
}
