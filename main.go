package main

import (
	"github.com/paulthom12345/elo_backend/models"
	"github.com/paulthom12345/elo_backend/server"
)

func main() {
	db, err := models.Connect()

	if err != nil {
		println("THERE WAS A PROBLEM CONNECTING TO THE DATABASE")
		panic(err.Error())
	}
	db.Migrate(true)

	server.StartServer(db)
	defer db.CloseDB()
}
