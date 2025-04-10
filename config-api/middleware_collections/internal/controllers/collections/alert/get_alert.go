package alert

import (
	"encoding/json"
	alertService "middleware_collections/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
)

// GetAlertHandler gère GET /alerts/{id}
func GetAlertHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.Context().Value("alertID").(string)
	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	result, err := alertService.GetAlertByID(id)
	if err != nil {
		http.Error(w, "Resource introuvable", http.StatusNotFound)
		return
	}

	if result == nil {
		http.Error(w, "Alerte introuvable", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
