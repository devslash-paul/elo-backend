package main

import (
	"github.com/paulthom12345/elo-backend/models"
	"github.com/paulthom12345/elo-backend/server"
	"github.com/paulthom12345/elo-backend/workers"
)

func main() {
	db, err := models.Connect()

	if err != nil {
		println("THERE WAS A PROBLEM CONNECTING TO THE `1")
		panic(err.Error())
	}
	db.Migrate(true)
	workers.BootStrap(db)

	server.StartServer(db)
	defer db.CloseDB()
}
