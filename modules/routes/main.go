package routes

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/leminhson2398/zipper/docs"
	"github.com/leminhson2398/zipper/modules/setting"
	custom_middle "github.com/leminhson2398/zipper/pkg/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// API registers all api routes this web app offers
func API() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Get("/token", TokenGenerator)
	router.With(custom_middle.ValidateToken).Post("/upload", ZipUploadHandler)

	// swagger API
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", setting.Port)),
	))

	return router
}
