package models

import (
	"errors"
	"time"
)

type League struct {
	ID           uint
	DeletedAt    *time.Time `xml:"-" json:"-"`
	CreatedAt    *time.Time `xml:"-" json:"-"`
	UpdatedAt    *time.Time `xml:"-" json:"-"`
	Name         string
	LeagueConfig []LeagueConfig `json:",omitempty" xml:",omitempty"`
}

type LeagueController struct {
	db DB
}

func NewLeagueController(db DB) *LeagueController {
	return &LeagueController{db}
}

func (ct *LeagueController) GetAllLeagues() *[]League {
	leagues := new([]League)
	ct.db.PreloadAllFor(leagues, "LeagueConfig")
	return leagues
}

// GetById will get a league by a specified league id
func (ct *LeagueController) GetById(id uint) (*League, error) {
	league := League{}

	println("The db looks like: ", ct)
	ct.db.First(&league, id)
	println(" Now the league looks like", league.ID)

	if league.ID == 0 {
		return nil, errors.New("No league has been found with that ID")
	}

	configs := []LeagueConfig{}
	ct.db.GetRelatedFor(&league, &configs)
	league.LeagueConfig = configs
	return &league, nil
}

func (ct *LeagueController) Create(league *League) error {
	ct.db.Create(league)
	return nil
}

func (ct *LeagueController) GetUserScores(league *League, users []uint) (int, int, error) {
	return 1, 1, nil
}
