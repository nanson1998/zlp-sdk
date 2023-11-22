package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func BuildMAC(key string, sep string, params ...interface{}) string {
	var str []string
	for _, p := range params {
		str = append(str, fmt.Sprint(p))
	}
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(strings.Join(str, sep)))

	return hex.EncodeToString(h.Sum(nil))
}
