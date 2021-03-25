package middleware

import (
	"net/http"

	"github.com/leminhson2398/zipper/pkg/logger"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil && err != http.ErrAbortHandler {

				logger.Logger.Error().Msgf("Got panic at path: %s, err: %v", r.URL.Path, err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
