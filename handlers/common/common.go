package common

import (
	"math/rand"
	"net/http"
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
