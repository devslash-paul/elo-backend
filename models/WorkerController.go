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

func NewWorkerController(db *ExportDB) WorkerController {
	return WorkerController{db}
}

func (ct *WorkerController) RegisterOrGetWorker(name string) Worker {
	worker := Worker{
		Name:               name,
		CurrentEventNumber: 0,
	}
	ct.db.Find("name = ?", worker.Name).Find(&worker)

	if worker.ID == 0 {
		ct.db.Create(worker)
	}

	return worker
}
