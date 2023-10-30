package middlewares

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RequestLogger(ctx *fiber.Ctx) error {
	start := time.Now()

	slog.Info(fmt.Sprintf("Receive %s %s", ctx.Method(), ctx.Path()),
		slog.String("method", ctx.Method()),
		slog.String("path", ctx.Path()),
		slog.String("url", ctx.OriginalURL()),
		slog.String("ip", ctx.IP()),
	)
	err := ctx.Next()

	latency := time.Since(start)

	slog.Info(fmt.Sprintf("Finish %s %s", ctx.Method(), ctx.Path()),
		slog.String("method", ctx.Method()),
		slog.String("path", ctx.Path()),
		slog.String("url", ctx.OriginalURL()),
		slog.String("ip", ctx.IP()),
		slog.Duration("latency", time.Duration(latency.Milliseconds())),
	)

	return err
}
