package server

import (
	"arville27/arv-toolkit/config"
	authService "arville27/arv-toolkit/modules/auth/service"
	splyrService "arville27/arv-toolkit/modules/splyr/service"
	"arville27/arv-toolkit/rest/middlewares"
	"arville27/arv-toolkit/rest/routes"
	"arville27/arv-toolkit/utils"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

	// Singleton
	httpClient := utils.NewHttpClient()
	spotifyLyricsService := splyrService.NewSpotifyLyricsService(
		httpClient,
		cfg.SpotifySPDCCookie,
		cfg.SpotifyTokenCachePath,
	)
	authService := authService.NewAuthService(cfg.AuthTokenSecret, cfg.ValidCredentials)

	server.app.Use(middlewares.RequestId)
	server.app.Use(middlewares.RequestLogger)
	apiV1 := server.app.Group("/api/v1")

	routes.SpotifyLyricsRoutes(apiV1.Group("/splyr", middlewares.Protected(cfg.AuthTokenSecret)), spotifyLyricsService)
	routes.AuthenticationRoutes(apiV1.Group("/auth"), authService)

	// Listen from a different goroutine
	go func() {
		if err := server.app.Listen(fmt.Sprintf("%s:%d", cfg.RestServerIp, cfg.RestServerPort)); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	server.app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	// Cleanup tasks go here
	fmt.Println("Fiber was successful shutdown.")
}
