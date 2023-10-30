package routes

import (
	"arville27/arv-toolkit/modules/splyr"
	"arville27/arv-toolkit/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func SpotifyLyricsRoutes(router fiber.Router, service splyr.SpotifyLyricsService) {
	router.Get("/", controllers.GetLyrics(service))
}
