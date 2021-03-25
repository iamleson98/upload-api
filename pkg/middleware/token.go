package middleware

import (
	"net/http"

	"github.com/leminhson2398/zipper/pkg/token"
)

const (
	authenHeader = "Authentication"
)

// ValidateTokenExist validates if a request contains authentication header token
func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header[authenHeader]
		if authHeader == nil || len(authHeader) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("You need to provide Authentication header to proceed"))
			return
		}
		tokenString := authHeader[0]
		// check if token still valid
		_, err := token.ValidateAccessToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("token was corrupted or expires"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
