package middleware

import (
	"strconv"
	"strings"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/repository"

	"github.com/gofiber/fiber/v3"
)

func Autentikasi() fiber.Handler {
	return func(c fiber.Ctx) error {
		header := c.Get("Authorization")

		if !strings.HasPrefix(header, "Bearer ") {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"Token autentikasi tidak ditemukan",
			)
		}

		tokenString := strings.TrimSpace(
			strings.TrimPrefix(header, "Bearer "),
		)

		konf := config.Muat()

		klaim, err := auth.PeriksaToken(
			tokenString,
			konf.JwtSecret,
		)
		if err != nil {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"Token tidak valid atau kedaluwarsa",
			)
		}

		// Periksa apakah token sudah masuk blacklist.
		terblokir, err := repository.CekTokenBlacklist(tokenString)
		if err != nil {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"Gagal memeriksa status token",
			)
		}

		if terblokir {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"Token sudah dicabut karena logout",
			)
		}

		idPengguna, err := strconv.ParseUint(
			klaim.Subject,
			10,
			32,
		)
		if err != nil {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"Identitas pengguna tidak valid",
			)
		}

		c.Locals("id_pengguna", uint(idPengguna))
		c.Locals("peran", klaim.Peran)
		c.Locals("token", tokenString)

		return c.Next()
	}
}

func HanyaPeran(peranDiizinkan ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		peran, ok := c.Locals("peran").(string)
		if !ok {
			return fiber.NewError(
				fiber.StatusForbidden,
				"Peran pengguna tidak ditemukan",
			)
		}

		for _, diizinkan := range peranDiizinkan {
			if peran == diizinkan {
				return c.Next()
			}
		}

		return fiber.NewError(
			fiber.StatusForbidden,
			"Anda tidak memiliki hak akses",
		)
	}
}
