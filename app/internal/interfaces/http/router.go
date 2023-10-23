package http

import (
	"EM-TASK/app/internal/infrastructure/logger"
	_ "EM-TASK/docs"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/swaggo/http-swagger/v2"
	"net/http"
	"os"
)

func (s *server) initRouter() http.Handler {
	r := s.router

	if os.Getenv("ENV") == "dev" {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	r.Post("/api/user", s.CreateUser)
	r.Get(`/api/user/{user_id}`, s.GetUser)
	r.Delete(`/api/user/{user_id}`, s.DeleteUser)
	r.Patch(`/api/user/{user_id}`, s.UpdateUser)
	r.Get(`/api/all_users`, s.GetAllUsers)
	r.Post(`/api/sort_users`, s.GetAllSortUsers)
	r.Handle("/", http.FileServer(http.Dir("public")))
	logger.Logger.Info().Msg("Route initialized")

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3005/swagger/doc.json"), // Замените URL на ваш путь к файлу swagger.json
	))
	logger.Logger.Info().Msg("Initialized task documentation")

	return r
}
