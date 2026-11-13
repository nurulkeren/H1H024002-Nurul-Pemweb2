package handler

import (
	"errors"
	"strings"

	"latihan-fiber-api/internal/auth"
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"
	"latihan-fiber-api/internal/validasi"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Registrasi pengguna baru
func Registrasi(c fiber.Ctx) error {
	var permintaan PermintaanRegistrasi

	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format data tidak valid")
	}

	galat := validasi.Periksa(permintaan)
	if len(galat) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data registrasi tidak valid",
			"galat":  galat,
		})
	}

	permintaan.Email = strings.TrimSpace(strings.ToLower(permintaan.Email))

	var penggunaLama model.Pengguna
	err := config.DB.Where("email = ?", permintaan.Email).First(&penggunaLama).Error

	if err == nil {
		return fiber.NewError(fiber.StatusConflict, "Email sudah terdaftar")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal memeriksa data pengguna")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(permintaan.KataSandi),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal memproses kata sandi")
	}

	pengguna := model.Pengguna{
		Nama:      permintaan.Nama,
		Email:     permintaan.Email,
		KataSandi: string(hash),
		Peran:     "mahasiswa",
	}

	if err := config.DB.Create(&pengguna).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menyimpan pengguna")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Registrasi berhasil",
		"data": fiber.Map{
			"id":    pengguna.ID,
			"nama":  pengguna.Nama,
			"email": pengguna.Email,
			"peran": pengguna.Peran,
		},
	})
}

// Login pengguna
func Login(c fiber.Ctx) error {
	var permintaan PermintaanLogin

	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format data tidak valid")
	}

	galat := validasi.Periksa(permintaan)
	if len(galat) > 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data login tidak valid",
			"galat":  galat,
		})
	}

	email := strings.TrimSpace(strings.ToLower(permintaan.Email))

	var pengguna model.Pengguna
	err := config.DB.Where("email = ?", email).First(&pengguna).Error

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Email atau kata sandi salah")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(pengguna.KataSandi),
		[]byte(permintaan.KataSandi),
	)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Email atau kata sandi salah")
	}

	konf := config.Muat()

	token, err := auth.TerbitkanToken(
		pengguna.ID,
		pengguna.Peran,
		konf.JwtSecret,
		konf.JwtTtlMenit,
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal membuat token")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Login berhasil",
		"data": fiber.Map{
			"token": token,
			"tipe":  "Bearer",
			"pengguna": fiber.Map{
				"id":    pengguna.ID,
				"nama":  pengguna.Nama,
				"email": pengguna.Email,
				"peran": pengguna.Peran,
			},
		},
	})
}

// Profil pengguna
func Profil(c fiber.Ctx) error {
	id := c.Locals("id_pengguna")

	var pengguna model.Pengguna

	if err := config.DB.First(&pengguna, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Pengguna tidak ditemukan")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data": fiber.Map{
			"id":    pengguna.ID,
			"nama":  pengguna.Nama,
			"email": pengguna.Email,
			"peran": pengguna.Peran,
		},
	})
}

// Refresh token JWT
func Segarkan(c fiber.Ctx) error {
	idPengguna, ok := c.Locals("id_pengguna").(uint)
	if !ok {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Identitas pengguna tidak valid",
		)
	}

	peran, ok := c.Locals("peran").(string)
	if !ok {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Peran pengguna tidak valid",
		)
	}

	konf := config.Muat()

	tokenBaru, err := auth.TerbitkanToken(
		idPengguna,
		peran,
		konf.JwtSecret,
		konf.JwtTtlMenit,
	)

	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal menerbitkan token baru",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Token berhasil diperbarui",
		"data": fiber.Map{
			"token": tokenBaru,
			"tipe":  "Bearer",
		},
	})
}

// Logout dan mencabut token JWT
func Logout(c fiber.Ctx) error {
	tokenString, ok := c.Locals("token").(string)
	if !ok || tokenString == "" {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Token autentikasi tidak ditemukan",
		)
	}

	// Ambil informasi kedaluwarsa token.
	konf := config.Muat()
	klaim, err := auth.PeriksaToken(tokenString, konf.JwtSecret)
	if err != nil {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Token tidak valid atau kedaluwarsa",
		)
	}

	if klaim.ExpiresAt == nil {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Masa berlaku token tidak ditemukan",
		)
	}

	// Simpan token ke blacklist.
	tokenBlacklist := model.TokenBlacklist{
		Token:     tokenString,
		ExpiresAt: klaim.ExpiresAt.Time,
	}

	err = repository.SimpanTokenBlacklist(&tokenBlacklist)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mencabut token",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Logout berhasil, token telah dicabut",
	})
}
