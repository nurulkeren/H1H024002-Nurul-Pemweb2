
package handler

import (
	"errors"
	"strconv"

	"latihan-fiber-api/internal/model"
	"latihan-fiber-api/internal/repository"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type PermintaanMataKuliah struct {
	Kode     string `json:"kode"`
	Nama     string `json:"nama"`
	SKS      int    `json:"sks"`
	Semester int    `json:"semester"`
}

func DaftarMataKuliah(c fiber.Ctx) error {
	daftar, err := repository.AmbilSemuaMataKuliah()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mata kuliah")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
	})
}

func DetailMataKuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID tidak valid")
	}

	matakuliah, err := repository.AmbilMataKuliahPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "Mata kuliah tidak ditemukan")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mata kuliah")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   matakuliah,
	})
}

func BuatMataKuliah(c fiber.Ctx) error {
	var permintaan PermintaanMataKuliah

	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format data tidak valid")
	}

	matakuliah := model.MataKuliah{
		Kode:     permintaan.Kode,
		Nama:     permintaan.Nama,
		SKS:      permintaan.SKS,
		Semester: permintaan.Semester,
	}

	if err := repository.SimpanMataKuliah(&matakuliah); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menyimpan mata kuliah")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil ditambahkan",
		"data":   matakuliah,
	})
}

func PerbaruiMataKuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID tidak valid")
	}

	matakuliah, err := repository.AmbilMataKuliahPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "Mata kuliah tidak ditemukan")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mata kuliah")
	}

	var permintaan PermintaanMataKuliah
	if err := c.Bind().Body(&permintaan); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Format data tidak valid")
	}

	if permintaan.Kode != "" {
		matakuliah.Kode = permintaan.Kode
	}
	if permintaan.Nama != "" {
		matakuliah.Nama = permintaan.Nama
	}
	if permintaan.SKS > 0 {
		matakuliah.SKS = permintaan.SKS
	}
	if permintaan.Semester > 0 {
		matakuliah.Semester = permintaan.Semester
	}

	if err := repository.PerbaruiMataKuliah(&matakuliah); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal memperbarui mata kuliah")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil diperbarui",
		"data":   matakuliah,
	})
}

func HapusMataKuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID tidak valid")
	}

	matakuliah, err := repository.AmbilMataKuliahPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusNotFound, "Mata kuliah tidak ditemukan")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal mengambil data mata kuliah")
	}

	if err := repository.HapusMataKuliah(&matakuliah); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Gagal menghapus mata kuliah")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Mata kuliah berhasil dihapus",
	})
}
