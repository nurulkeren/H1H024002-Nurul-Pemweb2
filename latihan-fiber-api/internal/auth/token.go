package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Klaim struct {
	Peran string `json:"peran"`
	jwt.RegisteredClaims
}

func TerbitkanToken(idPengguna uint, peran, rahasia string, ttlMenit int) (string, error) {
	waktuSekarang := time.Now()

	klaim := Klaim{
		Peran: peran,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(idPengguna), 10),
			IssuedAt:  jwt.NewNumericDate(waktuSekarang),
			ExpiresAt: jwt.NewNumericDate(waktuSekarang.Add(time.Duration(ttlMenit) * time.Minute)),
			Issuer:    "pemweb2-fiber-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, klaim)

	return token.SignedString([]byte(rahasia))
}

func PeriksaToken(nilaiToken, rahasia string) (*Klaim, error) {
	token, err := jwt.ParseWithClaims(
		nilaiToken,
		&Klaim{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("metode tanda tangan token tidak valid")
			}

			return []byte(rahasia), nil
		},
	)

	if err != nil {
		return nil, err
	}

	klaim, ok := token.Claims.(*Klaim)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return klaim, nil
}
