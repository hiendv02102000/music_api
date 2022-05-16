package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
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
func FormatStringSpace(s string) string {

	words := strings.Split(s, " ")
	res := ""
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		res += word + " "
	}
	res = strings.TrimSpace(res)
	return res
}
func ConverInterfaceToString(vals []interface{}) ([]string, error) {
	rs := make([]string, len(vals))
	for i, v := range vals {
		val, ok := v.(string)
		if !ok {
			return nil, errors.New("this is not string array")
		}
		rs[i] = val
	}
	return rs, nil
}
