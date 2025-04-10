package alert

import (
	"encoding/json"
	alertService "middleware_collections/internal/services/collections"
	"net/http"
)

// GET /alerts
func GetAlertsHandler(w http.ResponseWriter, r *http.Request) {
	alerts, err := alertService.GetAlerts()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des alertes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(alerts)
}
