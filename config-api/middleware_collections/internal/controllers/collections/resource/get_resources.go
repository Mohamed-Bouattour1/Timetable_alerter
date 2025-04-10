package resource

import (
	"encoding/json"
	resouceService "middleware_collections/internal/services/collections"
	"net/http"
)

// GetResourcesHandler godoc
// @Summary      Liste toutes les resources
// @Description  Retourne tous les emplois du temps configurés
// @Tags         resources
// @Produce      json
// @Success      200  {array}   models.Resource
// @Failure      500  {string}  string "Erreur serveur"
// @Router       /resources [get]
func GetResourcesHandler(w http.ResponseWriter, r *http.Request) {
	resources, err := resouceService.GetResources()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des resources", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resources)
}
