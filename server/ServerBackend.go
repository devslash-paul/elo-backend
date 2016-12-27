package server

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/paulthom12345/elo-backend/models"
)

type Registerable interface {
	register(container *restful.Container, endpoint string, db *models.ExportDB)
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
		endpoint: "/league",
	},
}

func NewServer(container *restful.Container, db *models.ExportDB) {
	for _, service := range services {
		service.rego.register(container, service.endpoint, db)
	}
}

func StartServer(db *models.ExportDB) {
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
