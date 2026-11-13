package router

import (
	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/handler"
	"latihan-fiber-api/internal/middleware"
)

func Daftarkan(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	// Route autentikasi publik
	api.Post("/auth/registrasi", handler.Registrasi)
	api.Post("/auth/login", handler.Login)

	// Route yang membutuhkan autentikasi JWT
	terlindungi := api.Group("", middleware.Autentikasi())

	// Route profil dan refresh token
	terlindungi.Get("/auth/profil", handler.Profil)
	terlindungi.Post("/auth/segarkan", handler.Segarkan)
	terlindungi.Post("/auth/logout", handler.Logout)
	// Route mahasiswa
	mahasiswa := terlindungi.Group("/mahasiswa")

	mahasiswa.Get("/", handler.DaftarMahasiswa)
	mahasiswa.Get("/:id", handler.DetailMahasiswa)

	mahasiswa.Post("/", middleware.HanyaPeran("admin", "dosen"), handler.BuatMahasiswa)
	mahasiswa.Put("/:id", middleware.HanyaPeran("admin", "dosen"), handler.PerbaruiMahasiswa)
	mahasiswa.Delete("/:id", middleware.HanyaPeran("admin"), handler.HapusMahasiswa)
	mahasiswa.Post("/:id/pulihkan", middleware.HanyaPeran("admin"), handler.PulihkanMahasiswa)

	// Route Mata Kuliah
	matakuliah := api.Group("/matakuliah")

	matakuliah.Get("/", handler.DaftarMataKuliah)
	matakuliah.Post("/", handler.BuatMataKuliah)
	matakuliah.Get("/:id", handler.DetailMataKuliah)
	matakuliah.Put("/:id", handler.PerbaruiMataKuliah)
	matakuliah.Delete("/:id", handler.HapusMataKuliah)

	// Route Program Studi
	api.Get("/program-studi", handler.DaftarProgramStudi)
}
