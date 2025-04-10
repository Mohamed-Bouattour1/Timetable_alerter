package alert

import (
	alertService "middleware_collections/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
)

// DeleteAlertHandler godoc
// @Summary      Supprime une alerte
// @Tags         alerts
// @Param        id   path      string  true  "ID de l'alerte"
// @Success      204  {string}  string "Suppression réussie"
// @Failure      400  {string}  string "ID invalide"
// @Router       /alerts/{id} [delete]
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
