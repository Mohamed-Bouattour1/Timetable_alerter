package internal

import (
	eventCtrl "timetable-api/internal/controllers/events"

	_ "timetable-api/api"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupRoutes configure toutes les routes
func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Routes pour les events (cours)
	r.Route("/events", func(r chi.Router) {
		r.Post("/", eventCtrl.PostEventHandler) // POST /events
		// On ajoutera GET, PUT, DELETE ici plus tard
	})

	// Route Swagger
	r.Handle("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return r
}
