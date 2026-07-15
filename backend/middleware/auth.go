package middleware

import (
	"net/http"
	"real-time-forum/backend/utils"
)

func GestOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if utils.LoggedIn(r) {
			utils.JSONError(w, http.StatusForbidden, "Aleardy logged in")
			return
		}
		next(w, r)
	}
}
