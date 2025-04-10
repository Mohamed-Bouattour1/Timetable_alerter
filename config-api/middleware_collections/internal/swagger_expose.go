package internal

import (
	alertsCtrl "middleware_collections/internal/controllers/collections/alert"
	resourcesCtrl "middleware_collections/internal/controllers/collections/resource"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() http.Handler {
	r := chi.NewRouter()

	// ROUTES RESOURCES
	r.Route("/resources", func(r chi.Router) {
		r.Get("/", resourcesCtrl.GetResourcesHandler)  // GET /resources
		r.Post("/", resourcesCtrl.PostResourceHandler) // POST /resources

		r.Route("/{id}", func(r chi.Router) {
			r.Use(resourcesCtrl.ResourceCtx)
			r.Get("/", resourcesCtrl.GetResourceHandler)       // GET /resources/{id}
			r.Put("/", resourcesCtrl.PutResourceHandler)       // PUT /resources/{id}
			r.Delete("/", resourcesCtrl.DeleteResourceHandler) // DELETE /resources/{id}
		})
	})

	// ROUTES ALERTS
	r.Route("/alerts", func(r chi.Router) {
		r.Get("/", alertsCtrl.GetAlertsHandler)  // GET /alerts
		r.Post("/", alertsCtrl.PostAlertHandler) // POST /alerts

		r.Route("/{id}", func(r chi.Router) {
			r.Use(alertsCtrl.AlertCtx)
			r.Get("/", alertsCtrl.GetAlertHandler)       // GET /alerts/{id}
			r.Put("/", alertsCtrl.PutAlertHandler)       // PUT /alerts/{id}
			r.Delete("/", alertsCtrl.DeleteAlertHandler) // DELETE /alerts/{id}
		})
	})

	return r
}
