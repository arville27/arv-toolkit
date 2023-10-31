package routes

import (
	"arville27/arv-toolkit/modules/auth"
	"arville27/arv-toolkit/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthenticationRoutes(router fiber.Router, service auth.AuthService) {
	router.Post("/", controllers.GenerateAccessToken(service))
}
