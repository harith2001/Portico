package gateway

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harith2001/portico/internal/config"
	"github.com/harith2001/portico/internal/services"
)

func NewApp(cfg config.Config) *fiber.App {
	app := fiber.New()

	// Global middlewares
	// app.Use(logger.Middleware())         // Logging
	// app.Use(auth.JWTMiddleware(cfg))     // JWT Authentication
	// app.Use(rate.RateLimiter(cfg))       // Per-client Rate Limiting

	// Routes
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Proxy route (catch-all) — secured behind JWT
	app.All("/proxy/*", func(c *fiber.Ctx) error {
		return services.HandleProxy(c)
	})

	return app
}
