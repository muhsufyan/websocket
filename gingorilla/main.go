package main

import (
	"log"

	"github.com/muhsufyan/websocket/routes"
)

func main() {

	r := routes.SetupRoutes()

	log.Println("listening on http://localhost:8040")
	r.Run(":8040")

}
