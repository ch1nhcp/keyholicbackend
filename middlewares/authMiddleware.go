package middlewares

import (
	"encoding/json"
	"net/http"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// cookie, _ := r.Cookie("jwt")
		// json.NewEncoder(w).Encode(cookie)
		header := r.Header.Get("Set-Token")

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}
		next.ServeHTTP(w, r)
	})
}
