package workers

import (
	"github.com/paulthom12345/elo_backend/models"
)

type Events []models.Event

// Update the GameEventWorker
// to the front of the EventQueue
func Update(db *models.ExportDB) {
	// var events Events
	// worker := db.GetWorkerFor("game_worker")
	// db.
	// 	Where("event_number > ?", worker.CurrentEventNumber).
	// 	Order("event_number asc").
	// 	Find(&events)
	// for _, event := range events {
	// 	processEvent(event)
	// }
}

func processEvent(event models.Event) {
	// so i know as the GameEventWorker that these events
	// are associated with the GameEvent table
	// var GameEvent gameEvent
	// db.First(&event, event.RelatedEventID)
}
