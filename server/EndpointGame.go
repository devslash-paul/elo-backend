package server

import (
	"strconv"

	restful "github.com/emicklei/go-restful"
	"github.com/paulthom12345/elo-backend/models"
)

type EndpointGame struct {
	db models.ExportDB
}

type endpointGamePutRequest struct {
	Winner uint64
	Loser  uint64
}

var gameController *models.GameController

func (ep *EndpointGame) register(container *restful.Container, endpoint string, db *models.ExportDB) {
	gameController = models.NewGameController(db)
	ws := new(restful.WebService)
	ws.
		Path(endpoint).
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("{league-id}/game/{game-id}").To(getGame).
		Doc("Get a game").
		Param(ws.PathParameter("league-id", "identifier of the league").DataType("int").Required(true)).
		Param(ws.PathParameter("game-id", "identifier of the game").DataType("int").Required(true)).
		Writes(models.Game{}))

	ws.Route(ws.PUT("{league-id}/game").To(createGame).
		Doc("Create a new game").
		Reads(models.Game{}).
		Writes(models.ModelValidation{}))

	ws.Route(ws.DELETE("{league-id}/game/{game-id}").To(deleteGame).
		Doc("Delete a single game").
		Param(ws.PathParameter("game-id", "Identifier of the game").DataType("int").Required(true)))

	container.Add(ws)
}

func createGame(req *restful.Request, resp *restful.Response) {
	game := models.Game{}
	var reqGame endpointGamePutRequest
	req.ReadEntity(&reqGame)

	id, _ := strconv.ParseUint(req.PathParameter("league-id"), 10, 64)

	println("REQ GAMES")
	println(reqGame.Loser)

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
