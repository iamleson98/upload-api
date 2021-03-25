package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	custom_middle "github.com/leminhson2398/zipper/pkg/middleware"
)

// API registers all api routes this web app offers
func API() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Get("/token", TokenGenerator)
	router.With(custom_middle.ValidateToken).Post("/upload", ZipUploadHandler)

	return router
}
