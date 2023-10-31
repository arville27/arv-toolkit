package controllers

import (
	"arville27/arv-toolkit/modules/auth"
	rest_model "arville27/arv-toolkit/rest/model"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GenerateAccessToken(s auth.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		type GenerateAccessTokenRequest struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		var requestBody *GenerateAccessTokenRequest
		err := ctx.BodyParser(&requestBody)
		if err != nil {
			slog.Error(
				"Failed to parse request body",
				slog.String("error", err.Error()),
			)
			return ctx.Status(http.StatusBadRequest).JSON(
				rest_model.RestErrorResponse("Failed to parse request body", "Bad request"),
			)
		}

		generatedTokenResponse, err := s.GenerateAccessToken(requestBody.Username, requestBody.Password)
		if err != nil {
			restError, statusCode := ResolveError(err)
			return ctx.Status(statusCode).JSON(restError)
		}

		responseBody := rest_model.RestSuccessResponse(generatedTokenResponse)
		return ctx.Status(200).JSON(responseBody)
	}
}
