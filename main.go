package main

import (
	"github.com/paulthom12345/elo-backend/models"
	"github.com/paulthom12345/elo-backend/server"
	"github.com/paulthom12345/elo-backend/workers"
)

func main() {
	db, err := models.Connect()

	if err != nil {
		println("There was a problem connecting to the database, program will now exit")
		panic(err.Error())
	}

	db.Migrate(true)
	workers.BootStrap(db)

	server.StartServer(db)
	defer db.CloseDB()
}
