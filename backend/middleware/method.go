package middleware

import (
	"net/http"
	"real-time-forum/backend/utils"
)

func Methode(method string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			utils.JSONError(w, 405, "Method not allowed")
			return
		}
		next(w, r)
	}
}
