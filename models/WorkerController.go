package models

import "time"

type Worker struct {
	ID                 uint
	DeletedAt          *time.Time `xml:"-" json:"-"`
	CreatedAt          *time.Time `xml:"-" json:"-"`
	UpdatedAt          *time.Time `xml:"-" json:"-"`
	CurrentEventNumber int
	Name               string
}

type WorkerController struct {
	db *ExportDB
}

func NewWorkerController(db *ExportDB) *WorkerController {
	return &WorkerController{db}
}

func (ct *WorkerController) RegisterOrGetWorker(name string) *Worker {
	worker := &Worker{
		Name:               name,
		CurrentEventNumber: 0,
	}

	ct.db.Where("name = ?", worker.Name).First(worker)

	if worker.ID == 0 {
		ct.db.Create(&worker)
	}

	return worker
}

func (ct *WorkerController) GetEventsForWorker(events []string, worker *Worker) []Event {
	output := new([]Event)
	ct.db.Where("EventName in (?) AND EventNumber > ?", events, worker.CurrentEventNumber).Find(&output)
	return *output
}
