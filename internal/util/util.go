package util

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"som/internal/config"

	"golang.org/x/crypto/bcrypt"
)

func NormalizeDate(date time.Time) (*time.Time, error) {
	cfg := config.Load()
	location, err := time.LoadLocation(cfg.Timezone)

	if err != nil {
		return nil, err
	}

	normalizedDate := time.Date(
		date.Year(),
		date.Month(),
		date.Day(),
		0, 0, 0, 0,
		location,
	).UTC()

	return &normalizedDate, nil
}

func GenerateMeetingCode() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b)[:8], nil
}

func HashPassword(password string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}

func ParseDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}
