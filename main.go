package main

import (
    "log"
    "laquinquenal/server"
    "laquinquenal/db"
)

func main() {
    db.ConnectDatabase()
    
    app := server.NewServer()
    if err := app.Start(); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
