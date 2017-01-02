package server

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/paulthom12345/elo-backend/models"
)

var gameController *models.GameController
var leagueController *models.LeagueController
var leagueConfigController *models.LeagueConfigController
var userController *models.UserController

type Registerable interface {
	register(ws *restful.WebService, endpoint string, db models.DB)
}

type Registerables []WrappedRegisterable

type WrappedRegisterable struct {
	rego     Registerable
	endpoint string
}

var services = Registerables{
	WrappedRegisterable{
		rego:     new(UserEndpoint),
		endpoint: "/users",
	},
	WrappedRegisterable{
		rego:     new(EndpointLeague),
		endpoint: "/leagues",
	},
	WrappedRegisterable{
		rego:     new(EndpointGame),
		endpoint: "/leagues",
	},
}

func NewServer(container *restful.Container, db models.DB) {
	wsPoints := make(map[string]*restful.WebService)
	gameController = models.NewGameController(db)
	leagueController = models.NewLeagueController(db)
	leagueConfigController = models.NewLeagueConfigController(db)
	userController = models.NewUserController(db)

	for _, service := range services {
		var wService *restful.WebService
		if ws, ok := wsPoints[service.endpoint]; ok {
			wService = ws
		} else {
			wService = &restful.WebService{}
			wsPoints[service.endpoint] = wService
			wService.
				Path(service.endpoint).
				Consumes(restful.MIME_JSON, restful.MIME_XML).
				Produces(restful.MIME_JSON, restful.MIME_XML)
		}

		service.rego.register(wService, service.endpoint, db)
	}

	for _, endpoint := range wsPoints {
		container.Add(endpoint)
	}
}

func StartServer(db models.DB) {
	wsContainer := restful.NewContainer()
	NewServer(wsContainer, db)

	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(),
		WebServicesUrl: "http://localhost:8888",
		ApiPath:        "/swagger.json",

		// Specify where the UI is located
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: "/users/pathompson/programming/swagger-ui-2.2.5/dist/",
	}
	swagger.RegisterSwaggerService(config, wsContainer)

	server := &http.Server{Addr: ":8888", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
