package services

import (
	"bytes"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var userBackends = map[string]string{
	// Example user to backend mapping (eventually loaded from DB)
	"user123": "http://localhost:9000", // user's backend
}

func HandleProxy(c *fiber.Ctx) error {
	userID := c.Locals("userID").(string) // Set by JWT middleware

	backendURL, ok := userBackends[userID]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No backend registered for user",
		})
	}

	// Build target URL
	targetURL := backendURL + c.Params("*") // e.g., /proxy/api/v1 -> http://backend/api/v1

	// Create new request
	req, err := http.NewRequest(c.Method(), targetURL, bytes.NewReader(c.Body()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "request creation failed"})
	}

	// Copy headers
	c.Request().Header.VisitAll(func(k, v []byte) {
		req.Header.Set(string(k), string(v))
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "backend error"})
	}
	defer resp.Body.Close()

	c.Status(resp.StatusCode)
	c.Response().SetBodyStream(resp.Body, -1)
	return nil
}
