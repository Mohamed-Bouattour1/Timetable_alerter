package resource

import (
	"encoding/json"
	resouceService "middleware_collections/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
)

// GetResourceHandler godoc
// @Summary      Récupère une resource par ID
// @Tags         resources
// @Produce      json
// @Param        id   path      string  true  "ID de la resource"
// @Success      200  {object}  models.Resource
// @Failure      404  {string}  string "Resource non trouvée"
// @Router       /resources/{id} [get]
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
