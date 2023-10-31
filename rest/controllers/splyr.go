package controllers

import (
	"arville27/arv-toolkit/modules/splyr"
	rest_model "arville27/arv-toolkit/rest/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetLyrics(s splyr.SpotifyLyricsService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		trackId := ctx.Query("trackId")

		if len(trackId) == 0 {
			return ctx.Status(http.StatusBadRequest).JSON(
				rest_model.RestErrorResponse(
					"Missing required query parameter 'trackId'",
					"Missing Required Field",
				),
			)
		}

		lyrics, err := s.GetLyrics(trackId)
		if err != nil {
			restError, statusCode := ResolveError(err)
			return ctx.Status(statusCode).JSON(restError)
		}

		responseBody := rest_model.RestSuccessResponse(lyrics)
		return ctx.Status(200).JSON(responseBody)
	}
}
