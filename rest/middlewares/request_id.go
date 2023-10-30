package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

const RequestIdContextKey = "requestId"

func RequestId(ctx *fiber.Ctx) error {
	// Get id from request, else we generate one
	requestId := utils.UUID()

	// Set new id to response header
	ctx.Set(fiber.HeaderXRequestID, requestId)

	// Add the request ID to locals
	ctx.Locals(RequestIdContextKey, requestId)

	// Continue stack
	return ctx.Next()
}
