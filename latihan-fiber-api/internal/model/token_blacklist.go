package model

import "time"

type TokenBlacklist struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Token     string    `gorm:"type:text;not null" json:"-"`
	ExpiresAt time.Time `gorm:"not null;index" json:"kedaluwarsa_pada"`
	CreatedAt time.Time `json:"dibuat_pada"`
}
