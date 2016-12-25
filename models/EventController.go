package models

import "time"

type Event struct {
	ID             uint
	DeletedAt      *time.Time `xml:"-" json:"-"`
	CreatedAt      *time.Time `xml:"-" json:"-"`
	UpdatedAt      *time.Time `xml:"-" json:"-"`
	EventNumber    uint64     `gorm:"AUTO_INCREMENT"`
	EventName      string
	WorkerID       int
	RelatedEventID uint // This works by a worker joining to a table
	RelatedTable   string
}

type EventController struct {
	db *ExportDB
}

func NewEventController(db *ExportDB) *EventController {
	return &EventController{db}
}

func (ct *EventController) CreateEvent(e *Event) {
	ct.db.Create(e)
}

func (ct *EventController) GetEventsForWorker(lastRead uint, events []string) []Event {
	output := new([]Event)
	ct.db.Where("EventName in (?) AND EventNumber > ?", events, lastRead).Find(&output)
	return *output
}
