package middlewares

import (
	"encoding/json"
	"finalbackend/util"
	"net/http"
)

func JwtVerify(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("admin")
		if cookie == nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}
		id, _ := util.ParseJwt(cookie.Value)
		// json.NewEncoder(w).Encode(cookie)
		// header := r.Header.Get("Set-Token")
		if id == "0" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("khong du quyen")
			return
		}
		next.ServeHTTP(w, r)
	})
}
