package main

import (
	"arville27/arv-toolkit/config"
	"arville27/arv-toolkit/logger"
	"arville27/arv-toolkit/rest/server"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfigFromEnv()
	logger.ConfigureLogger()

	// REST API Server
	fiberApp := fiber.New(fiber.Config{
		AppName: "Arv Toolkit",
	})
	restServer := server.NewRestServer(fiberApp, cfg)
	restServer.Start()
}
