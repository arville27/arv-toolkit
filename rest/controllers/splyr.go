package controllers

import (
	"arville27/arv-toolkit/modules/splyr"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetLyrics(s splyr.SpotifyLyricsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		trackId := c.Query("trackId")

		if len(trackId) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"data":  nil,
				"error": "Missing required query parameter 'trackId'",
			})
		}

		lyrics, err := s.GetLyrics(trackId)
		if err != nil {
			var splyrError *splyr.SplyrError
			if errors.As(err, &splyrError) {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"data":  nil,
					"error": err.Error(),
				})
			}
			return c.Status(500).JSON(&fiber.Map{
				"data":  nil,
				"error": err.Error(),
			})
		}

		return c.Status(200).JSON(&fiber.Map{
			"data":  lyrics,
			"error": nil,
		})
	}
}
