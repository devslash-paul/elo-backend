package workers

import (
	"github.com/paulthom12345/elo_backend/models"
)

type Events []models.Event

type ELOWorker struct {
	models.Worker
}

func NewEloWorker(db models.Worker) *ELOWorker {
	return &ELOWorker{db}
}

func (work *ELOWorker) Update() {

}
