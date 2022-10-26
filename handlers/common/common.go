package common

import (
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

func HandleError(err error, w http.ResponseWriter, r *http.Request, status int) {
	if status == 0 {
		status = http.StatusInternalServerError
	}
	http.Error(w, "internal server error: "+err.Error(), status)
}

var randomCharacterList = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomCharacters(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = randomCharacterList[rand.Intn(len(randomCharacterList))]
	}
	return string(b)
}

func GetQueryParam(key string, url *url.URL) string {
	parts := strings.Split(url.RawQuery, "&")
	for _, part := range parts {
		param := strings.Split(part, "=")
		if param[0] == key {
			return param[1]
		}
	}
	return ""
}
