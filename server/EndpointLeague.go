package server

import (
	"fmt"
	"net/http"
	"strconv"

	restful "github.com/emicklei/go-restful"
	"github.com/paulthom12345/elo-backend/models"
)

type EndpointLeague struct {
	db *models.DB
}

func (u *EndpointLeague) register(ws *restful.WebService, endpoint string, db models.DB) {
	leagueController = models.NewLeagueController(db)
	leagueConfigController = models.NewLeagueConfigController(db)

	ws.Route(ws.GET("/").To(getLeagues))

	ws.Route(ws.GET("/{league-id}").To(getLeague).
		Doc("Get a league").
		Param(ws.PathParameter("league-id", "identifier of the league").DataType("int").Required(true)).
		Writes(&models.League{}))

	ws.Route(ws.PUT("/").To(createLeague).
		Doc("Create a new League with the default settings.").
		Operation("updateUser").
		Reads(models.League{})) // from the request

	ws.Route(ws.PUT("/{league-id}/config").To(createConfig).
		Doc("Create a new league config. Fail if setting exists already").
		Reads(models.LeagueConfig{}).
		Param(ws.PathParameter("league-id", "identifier of the league").DataType("int").Required(true)))
}

func getLeagues(req *restful.Request, resp *restful.Response) {
	leagues := leagueController.GetAllLeagues()

	resp.WriteEntity(leagues)
}

func createConfig(req *restful.Request, resp *restful.Response) {
	id, _ := strconv.ParseUint(req.PathParameter("league-id"), 10, 64)

	var reqLeagueConfig models.LeagueConfig
	req.ReadEntity(&reqLeagueConfig)

	reqLeagueConfig.LeagueID = id

	leagueConfigController.Create(&reqLeagueConfig)
	resp.WriteEntity(reqLeagueConfig)
}

func createLeague(req *restful.Request, resp *restful.Response) {
	var reqLeague models.League
	err := req.ReadEntity(&reqLeague)
	reqLeague.ID = 0 // Can i skip this from being read as an entity?
	if err != nil {
		fmt.Println(err.Error())
		resp.WriteHeaderAndEntity(http.StatusBadRequest, err.Error())
		return
	}
	leagueController.Create(&reqLeague)
	resp.WriteEntity(reqLeague)
}

func getLeague(req *restful.Request, resp *restful.Response) {
	id, _ := strconv.ParseUint(req.PathParameter("league-id"), 10, 64)
	league, err := leagueController.GetById(uint(id))

	if err != nil || league.DeletedAt != nil {
		resp.WriteErrorString(http.StatusNotFound, "League could not be found")
		return
	}

	resp.WriteEntity(league)
}
