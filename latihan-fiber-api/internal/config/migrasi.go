
package config

import (
	"log"

	"latihan-fiber-api/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Migrasi tabel database
func Migrasi() {
	err := DB.AutoMigrate(
		&model.ProgramStudi{},
		&model.Mahasiswa{},
		&model.MataKuliah{},
		&model.Pengguna{},
		&model.TokenBlacklist{},
	)

	if err != nil {
		log.Fatal("Gagal menjalankan migrasi: ", err)
	}

	log.Println("Migrasi selesai")
}

// Membuat data awal program studi
func DataAwal() {
	var jumlah int64

	DB.Model(&model.ProgramStudi{}).Count(&jumlah)

	if jumlah > 0 {
		return
	}

	daftar := []model.ProgramStudi{
		{Kode: "TK", Nama: "Teknik Komputer", Jenjang: "S1"},
		{Kode: "IF", Nama: "Informatika", Jenjang: "S1"},
		{Kode: "TE", Nama: "Teknik Elektro", Jenjang: "S1"},
	}

	err := DB.Create(&daftar).Error
	if err != nil {
		log.Fatal("Gagal membuat data awal program studi: ", err)
	}

	log.Println("Data awal program studi dibuat")
}

// Membuat akun admin awal
func BuatAdminAwal() {
	email := "admin@unsoed.ac.id"

	var admin model.Pengguna

	err := DB.Where("email = ?", email).First(&admin).Error

	if err == nil {
		log.Println("Akun admin sudah tersedia")
		return
	}

	if err != gorm.ErrRecordNotFound {
		log.Fatal("Gagal memeriksa akun admin: ", err)
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte("Admin12345!"),
		bcrypt.DefaultCost,
	)

	if err != nil {
		log.Fatal("Gagal membuat password admin: ", err)
	}

	nama := "Administrator"
	peran := "admin"
	kataSandi := string(hash)

	admin = model.Pengguna{
		Nama:      nama,
		Email:     email,
		KataSandi: kataSandi,
		Peran:     peran,
	}

	if err := DB.Create(&admin).Error; err != nil {
		log.Fatal("Gagal membuat akun admin: ", err)
	}

	log.Println("Akun admin berhasil dibuat")
}
