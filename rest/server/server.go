package server

import (
	"arville27/arv-toolkit/config"
	splyrService "arville27/arv-toolkit/modules/splyr/service"
	"arville27/arv-toolkit/rest/middlewares"
	"arville27/arv-toolkit/rest/routes"
	"arville27/arv-toolkit/utils"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RestServer interface {
	Start()
}

type restServer struct {
	app *fiber.App
	cfg *config.Config
}

func NewRestServer(app *fiber.App, cfg *config.Config) RestServer {
	return restServer{app, cfg}
}

func (server restServer) Start() {
	cfg := server.cfg

	httpClient := utils.NewHttpClient()
	spotifyLyricsService := splyrService.NewSpotifyLyricsService(
		httpClient,
		cfg.SpotifySPDCCookie,
		cfg.SpotifyTokenCachePath,
	)

	server.app.Use(middlewares.RequestId)
	server.app.Use(middlewares.RequestLogger)
	apiV1 := server.app.Group("/api/v1")

	routes.SpotifyLyricsRoutes(apiV1.Group("/splyr"), spotifyLyricsService)

	log.Fatal(server.app.Listen(fmt.Sprintf("%s:%d", cfg.RestServerIp, cfg.RestServerPort)))
}
