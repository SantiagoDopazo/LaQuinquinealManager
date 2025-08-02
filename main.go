package main

import (
    "log"
    "laquinquenal/config"
    "laquinquenal/server"
)

func main() {
    config.ConnectDatabase()
    
    app := server.NewServer()
    if err := app.Start(); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
