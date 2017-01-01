package models

import "time"

type Worker struct {
	ID                 uint
	DeletedAt          *time.Time `xml:"-" json:"-"`
	CreatedAt          *time.Time `xml:"-" json:"-"`
	UpdatedAt          *time.Time `xml:"-" json:"-"`
	CurrentEventNumber int
	Name               string `gorm:"type:varchar(100);unique"`
}

type WorkerController struct {
	db DB
}

func NewWorkerController(db DB) *WorkerController {
	return &WorkerController{db}
}

func (ct *WorkerController) RegisterOrGetWorker(name string) *Worker {
	worker := &Worker{
		Name:               name,
		CurrentEventNumber: 0,
	}

	ct.db.Where(worker, "name = ?", worker.Name)

	if worker.ID == 0 {
		ct.db.Create(&worker)
	}

	return worker
}

func (ct *WorkerController) GetEventsForWorker(events []string, worker *Worker) []Event {
	output := new([]Event)
	ct.db.Where(output, "EventName in (?) AND EventNumber > ?", events, worker.CurrentEventNumber)
	return *output
}
