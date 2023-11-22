package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

var (
	PROD      = "https://openapi.zalopay.vn"
	SB        = "https://sb-openapi.zalopay.vn"
	DefaultId = 0
)

func init() {
	DefaultId = 1
}

func GetAppTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func GetTimeString(apptime int64) string {
	t := time.Unix(apptime/1000, 0)
	return fmt.Sprintf("%02d%02d%02d", t.Year()%100, int(t.Month()), t.Day())
}

// yymmdd_appid_xxxxxxxxxx
func GetMRefundId(appid string, apptime int64) string {
	if DefaultId >= 100000 {
		DefaultId = 0
	}

	DefaultId += 1
	t := time.Now().Format("150405")
	return fmt.Sprintf("%v_%v_%v%05d", GetTimeString(apptime), appid, t, DefaultId)
}

func BuildMAC(key string, sep string, params ...interface{}) string {
	var str []string
	for _, p := range params {
		str = append(str, fmt.Sprint(p))
	}
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(strings.Join(str, sep)))

	return hex.EncodeToString(h.Sum(nil))
}
