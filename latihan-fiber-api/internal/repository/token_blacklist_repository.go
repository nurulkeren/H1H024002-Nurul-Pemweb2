package repository

import (
	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func SimpanTokenBlacklist(token *model.TokenBlacklist) error {
	return config.DB.Create(token).Error
}

func CekTokenBlacklist(nilaiToken string) (bool, error) {
	var jumlah int64

	err := config.DB.Model(&model.TokenBlacklist{}).
		Where("token = ?", nilaiToken).
		Count(&jumlah).Error

	return jumlah > 0, err
}
