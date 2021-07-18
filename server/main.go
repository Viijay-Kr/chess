package main

import (
	"chess/game/db"
	"chess/game/room"
	"log"
	"net/http"
)

func main() {
	gameServer := room.CreateConnection()

	go func() {
		if err := gameServer.Serve(); err != nil {
			log.Fatalf("failed to start socket server %s\n", err)
		}
	}()
	defer gameServer.Close()
	db.SetupDatabase()

	http.Handle("/", gameServer)

	log.Println("Serving at localhost:8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
