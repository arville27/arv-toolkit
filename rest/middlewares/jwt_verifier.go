package middlewares

import (
	rest_model "arville27/arv-toolkit/rest/model"
	"log/slog"
	"net/http"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func Protected(signingKeyKey string) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(signingKeyKey)},
		ErrorHandler: jwtError,
	})
}

func jwtError(ctx *fiber.Ctx, err error) error {
	slog.Info(
		"From token middleware",
		slog.String("error", err.Error()),
	)
	if err.Error() == "Missing or malformed JWT" {
		return ctx.Status(http.StatusBadRequest).JSON(
			rest_model.RestErrorResponse(
				"Bad request",
				"Missing or malformed token",
			),
		)
	} else {
		return ctx.Status(http.StatusUnauthorized).JSON(
			rest_model.RestErrorResponse(
				"Unauthorized",
				"Invalid or expired token",
			),
		)
	}
}
