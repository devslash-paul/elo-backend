package workers

import "github.com/paulthom12345/elo-backend/models"

type RunnableWorker interface {
	Update()
}

type Workers []RunnableWorker

type WorkerManager struct {
	workers Workers
}

type WorkerManagerIn interface {
	Touch()
}

var workerRegistry Workers

func BootStrap(db models.DB) *WorkerManager {
	workerCT := models.NewWorkerController(db)

	workerRegistry = Workers{
		NewEloWorker(workerCT.RegisterOrGetWorker("ELOWorker"), db),
	}

	return &WorkerManager{workerRegistry}
}

func (wm *WorkerManager) AfterEvent() {
	for _, worker := range wm.workers {
		worker.Update()
	}
}
