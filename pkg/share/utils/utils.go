package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"time"
)

const time_layout = "2006-01-02 15:04:05"

func EncryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(h[:])
}
func FormatTime(time time.Time) string {
	return time.Format(time_layout)
}
