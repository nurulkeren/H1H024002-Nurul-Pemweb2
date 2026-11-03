
package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

func KunciApi() fiber.Handler {
	return func(c fiber.Ctx) error {
		apiKey := c.Get("X-Api-Key")
		apiKeyBenar := os.Getenv("API_KEY")

		if apiKeyBenar == "" || apiKey != apiKeyBenar {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"API Key tidak valid",
			)
		}

		return c.Next()
	}
}
