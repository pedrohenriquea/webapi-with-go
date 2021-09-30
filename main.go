package main

import (
	"github.com/pedrohenriquea/webapi-with-go/database"
	"github.com/pedrohenriquea/webapi-with-go/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()
}
