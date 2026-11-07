package router

import (
	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/handler"
)

func Daftarkan(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"sukses": true,
			"pesan":  "API Fiber aktif",
		})
	})

	// Route Mahasiswa
	mahasiswa := api.Group("/mahasiswa")

	mahasiswa.Get("/", handler.DaftarMahasiswa)
	mahasiswa.Post("/", handler.BuatMahasiswa)
	mahasiswa.Get("/:id", handler.DetailMahasiswa)
	mahasiswa.Put("/:id", handler.PerbaruiMahasiswa)
	mahasiswa.Delete("/:id", handler.HapusMahasiswa)
	mahasiswa.Post("/:id/pulihkan", handler.PulihkanMahasiswa)

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
