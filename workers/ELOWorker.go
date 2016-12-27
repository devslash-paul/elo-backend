package workers

import (
	"github.com/paulthom12345/elo-backend/models"
)

type Events []models.Event

type ELOWorker struct {
	worker  *models.Worker
	workC   *models.WorkerController
	userC   *models.UserController
	leagueC *models.LeagueController
	gameC   *models.GameController
}

func NewEloWorker(worker *models.Worker, db *models.ExportDB) *ELOWorker {
	return &ELOWorker{
		worker,
		models.NewWorkerController(db),
		models.NewUserController(db),
		models.NewLeagueController(db),
		models.NewGameController(db)}
}

// ELO worker updates by getting every 'game' event in its queue
func (work *ELOWorker) Update() {
	parseEvents := work.workC.GetEventsForWorker([]string{"game"}, work.worker)
	for _, event := range parseEvents {
		work.registerGame(event)
	}
}

func (work *ELOWorker) registerGame(event models.Event) {
	game, err := work.gameC.GetGameForEvent(event)
	league, err := work.leagueC.GetById(uint(game.LeagueID))

	if err != nil {
		panic(err)
	}

	winnerScore, loserScore, err := work.leagueC.GetUserScores(league, []uint{game.WinnerID, game.LoserID})
	println(winnerScore)
	println(loserScore)
	println(err)
}
