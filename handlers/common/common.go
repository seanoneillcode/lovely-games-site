package common

import "net/http"

func HandleError(err error, w http.ResponseWriter, r *http.Request) {
	http.Error(w, "internal server error: "+err.Error(), http.StatusInternalServerError)
}
