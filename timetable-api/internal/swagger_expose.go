// @title           Timetable API
// @version         1.0
// @description     Microservice de gestion des événements de l'emploi du temps
// @host      localhost:8080
// @BasePath  /

package internal

import (
	eventCtrl "timetable-api/internal/controllers/events"

	_ "timetable-api/api"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// configurer toutes les routes
func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Routes pour les events
	r.Route("/events", func(r chi.Router) {
		r.Get("/", eventCtrl.GetAllEventsHandler) // GET /events
	})

	// Route Swagger
	r.Handle("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	return r
}
