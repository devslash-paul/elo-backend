package server

import (
	"strconv"

	restful "github.com/emicklei/go-restful"
	"github.com/paulthom12345/elo-backend/models"
)

type EndpointGame struct {
	db models.DB
}

type endpointGamePutRequest struct {
	Winner uint64
	Loser  uint64
}

func (ep *EndpointGame) register(ws *restful.WebService, endpoint string, db models.DB) {
	gameController = models.NewGameController(db)
	leagueController = models.NewLeagueController(db)
	userController = models.NewUserController(db)

	ws.Route(ws.GET("/{league-id}/game/{game-id}").To(getGame).
		Doc("Get a game").
		Param(ws.PathParameter("league-id", "identifier of the league").DataType("int").Required(true)).
		Param(ws.PathParameter("game-id", "identifier of the game").DataType("int").Required(true)).
		Writes(models.Game{}))

	ws.Route(ws.PUT("/{league-id}/game").To(createGame).
		Doc("Create a new game").
		Reads(models.Game{}).
		Writes(models.ModelValidation{}))

	ws.Route(ws.DELETE("/{league-id}/game/{game-id}").To(deleteGame).
		Doc("Delete a single game").
		Param(ws.PathParameter("game-id", "Identifier of the game").DataType("int").Required(true)))
}

func createGame(req *restful.Request, resp *restful.Response) {
	game := models.Game{}

	var reqGame endpointGamePutRequest
	err := req.ReadEntity(&reqGame)
	if err != nil {
		panic(err)
	}

	id, _ := strconv.ParseUint(req.PathParameter("league-id"), 10, 64)

	league, _ := leagueController.GetById(uint(id))
	winner, _ := userController.GetUserByID(reqGame.Winner)
	loser, _ := userController.GetUserByID(reqGame.Loser)

	game.League = *league
	game.Winner = *winner
	game.Loser = *loser

	gameController.Create(&game)
	resp.WriteEntity(&game)
}

func getGame(req *restful.Request, resp *restful.Response) {
	leagueid, _ := strconv.ParseUint(req.PathParameter("league-id"), 10, 64)
	gameid, _ := strconv.ParseUint(req.PathParameter("game-id"), 10, 64)

	league, _ := leagueController.GetById(uint(leagueid))
	game, _ := gameController.GetById(league, gameid)

	resp.WriteEntity(game)
}

func deleteGame(req *restful.Request, resp *restful.Response) {

}
