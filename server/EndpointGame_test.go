package server

import (
	"bytes"
	"net/http"
	"testing"

	"encoding/json"

	restful "github.com/emicklei/go-restful"
	"github.com/paulthom12345/elo-backend/mocks"
	"github.com/paulthom12345/elo-backend/models"
	. "github.com/stretchr/testify/mock"
)

var server *http.Server
var db *mocks.DB

var currentGame *models.Game

func TestMain(t *testing.T) {
	epg := new(EndpointGame)
	wsContainer := restful.NewContainer()
	db = new(mocks.DB)
	db.On("GetRelatedFor", AnythingOfType("*models.League"), Anything).Run(func(a Arguments) {
	})
	db.On("Create", AnythingOfType("*models.Game")).Run(func(a Arguments) {
		cgame, ok := (a.Get(0)).(*models.Game)
		if !ok {
			panic(ok)
		}

		currentGame = cgame
		currentGame.ID = 1
		println("Creatng with name: ", currentGame.Winner.FullName)
	})
	db.On("First", AnythingOfType("*models.League"), Anything).Run(func(a Arguments) {
		league, ok := (a.Get(0)).(*models.League)
		if !ok {
			panic(ok)
		}
		league.ID = 1
	})
	db.On("First", AnythingOfType("*models.User"), Anything).Run(func(a Arguments) {
		user, ok := (a.Get(0)).(*models.User)
		if !ok {
			panic(ok)
		}
		user.ID = 1
		user.FullName = "Fish"
	})
	db.On("PrimaryWithCondidtion", AnythingOfType("*models.Game"), uint64(1), "League = ?", Anything).Run(func(a Arguments) {
		i, _ := a.Get(0).(*models.Game)
		i.ID = currentGame.ID
		i.League = currentGame.League
		i.Winner = currentGame.Winner
		i.WinnerID = currentGame.WinnerID
	})
	epg.register(wsContainer, "/league", db)
	server = &http.Server{Addr: ":8888", Handler: wsContainer}
	go server.ListenAndServe()
}

func TestGameIsCreated(t *testing.T) {
	urlstr := "http://localhost:8888/league/1/game"
	client := &http.Client{}
	jsonStr := []byte(`{
		"Winner": 1,
		"League": 1
		}`)
	request, err := http.NewRequest("PUT", urlstr, bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Error when attempting to create the game")
	}

	// Check the game
	if currentGame.League.ID != 1 {
		t.Fatalf("League was not set", currentGame)
	}

	// try get that game again
	gettsr := "http://localhost:8888/league/1/game/1"
	resp, _ = http.Get(gettsr)
	if resp.StatusCode != 200 {
		t.Errorf("Status code should be 200, was %d", resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	var dat models.Game

	if err := json.Unmarshal(buf.Bytes(), &dat); err != nil {
		panic(err)
	}

	if dat.Winner.ID != 1 {
		t.Errorf("Winner was not 1")
	}
}
