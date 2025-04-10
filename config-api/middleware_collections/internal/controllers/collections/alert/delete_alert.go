package alert

import (
	alertService "middleware_collections/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
)

// DELETE /alerts/{id}
func DeleteAlertHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.Context().Value("alertID").(string)
	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	err = alertService.DeleteAlert(id)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
