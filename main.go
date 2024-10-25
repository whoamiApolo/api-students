package main

import (
	"github.com/whoamiApolo/api-students/api"
	"log"
)

func main() {
	server := api.NewServer()
	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
