package server

import (
	"fmt"
	"net/http"
	"strconv"

	restful "github.com/emicklei/go-restful"
	"github.com/paulthom12345/elo-backend/models"
)

type UserEndpoint struct {
	endpoint string
}

type errormsg struct {
	Message string
}

var userController *models.UserController

func (u *UserEndpoint) register(container *restful.Container, endpoint string, db *models.ExportDB) {
	userController = models.NewUserController(db)
	ws := new(restful.WebService)

	ws.
		Path(endpoint).
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{user-id}").To(getUser).
		Doc("Get a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("int").Required(true)).
		Writes(models.User{}))

	ws.Route(ws.PUT("/").To(createUser).
		Doc("Create a new User").
		Reads(models.User{}).
		Writes(models.ModelValidation{}))

	ws.Route(ws.DELETE("/{user-id}").To(deleteUser).
		Doc("Delete a single user").
		Param(ws.PathParameter("user-id", "Identifier of the user").DataType("int").Required(true)))

	container.Add(ws)
}

func createUser(req *restful.Request, resp *restful.Response) {
	var reqUser models.User
	err := req.ReadEntity(&reqUser)

	validation := reqUser.Validate()

	if validation != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, validation)
		return
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	userController.Create(&reqUser)
	resp.WriteEntity(reqUser)
}

func deleteUser(req *restful.Request, resp *restful.Response) {
	id, _ := strconv.ParseUint(req.PathParameter("user-id"), 10, 64)
	user, err := userController.GetUserByID(id)

	if err != nil || user.DeletedAt != nil {
		resp.WriteHeaderAndEntity(http.StatusNotFound, errormsg{Message: "User could not be found"})
	} else {
		userController.Delete(user)
		resp.WriteHeader(http.StatusAccepted)
	}
}

func getUser(req *restful.Request, resp *restful.Response) {
	id, _ := strconv.ParseUint(req.PathParameter("user-id"), 10, 64)
	user, err := userController.GetUserByID(id)

	if err != nil || user.DeletedAt != nil {
		resp.WriteHeaderAndEntity(http.StatusNotFound, errormsg{Message: "User could not be found"})
	} else {
		resp.WriteEntity(user)
	}
}
