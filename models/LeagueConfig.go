package models

import (
	"errors"
	"time"
)

type LeagueConfig struct {
	ID          uint
	DeletedAt   *time.Time `xml:"-" json:"-"`
	CreatedAt   *time.Time `xml:"-" json:"-"`
	UpdatedAt   *time.Time `xml:"-" json:"-"`
	ConfigName  string
	ConfigValue string

	LeagueID uint64
}

type LeagueConfigController struct {
	db DB
}

func NewLeagueConfigController(db DB) *LeagueConfigController {
	return &LeagueConfigController{db}
}

func (ct *LeagueConfigController) GetById(id uint64) (*LeagueConfig, error) {
	config := LeagueConfig{}
	ct.db.First(&config, id)

	if config.ID == 0 {
		return nil, errors.New("No league has been found with that ID")
	}

	return &config, nil
}

func (ct *LeagueConfigController) Create(config *LeagueConfig) error {
	ct.db.Create(config)
	return nil
}
