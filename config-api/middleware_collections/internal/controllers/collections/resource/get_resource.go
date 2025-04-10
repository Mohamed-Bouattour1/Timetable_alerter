package resource

import (
	"encoding/json"
	resouceService "middleware_collections/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
)

// GET /resources/{id}
func GetResourceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.Context().Value("resourceID").(string)
	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	result, err := resouceService.GetResourceByID(id)
	if err != nil {
		http.Error(w, "Resource introuvable", http.StatusNotFound)
		return
	}

	if result == nil {
		http.Error(w, "Resource introuvable", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
