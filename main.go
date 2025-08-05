package main

import (
	"laquinquenal/db"
	"laquinquenal/server"
	"log"
)

func main() {
	db.ConnectDatabase()

	app := server.NewServer()
	if err := app.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
