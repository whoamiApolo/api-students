package main

import (
	"github.com/rs/zerolog/log"
	"github.com/whoamiApolo/api-students/api"
)

func main() {
	server := api.NewServer()
	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msgf("Failed to start server: %s", err.Error())
	}
}
