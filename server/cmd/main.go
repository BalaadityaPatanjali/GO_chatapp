package main

import (
	"log"

	"github.com/BalaadityaPatanjali/GO_chatapp/db"
	"github.com/BalaadityaPatanjali/GO_chatapp/internal/user"
	"github.com/BalaadityaPatanjali/GO_chatapp/internal/ws"
	"github.com/BalaadityaPatanjali/GO_chatapp/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	// instantiate web socket hub
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	go hub.Run()


	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
