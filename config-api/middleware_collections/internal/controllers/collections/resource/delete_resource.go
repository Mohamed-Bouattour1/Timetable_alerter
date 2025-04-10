package resource

import (
	"middleware_collections/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
)

// DeleteResourceHandler godoc
// @Summary      Supprime une resource
// @Tags         resources
// @Param        id   path      string  true  "ID de la resource"
// @Success      204  {string}  string "Suppression réussie"
// @Failure      400  {string}  string "ID invalide"
// @Router       /resources/{id} [delete]
func DeleteResourceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.Context().Value("resourceID").(string)
	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	err = collections.DeleteResource(id)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
