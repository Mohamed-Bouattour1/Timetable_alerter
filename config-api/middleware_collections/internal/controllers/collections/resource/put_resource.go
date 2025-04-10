package resource

import (
	"encoding/json"
	"middleware_collections/internal/models"
	"middleware_collections/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
)

// PutResourceHandler godoc
// @Summary      Met à jour une resource
// @Tags         resources
// @Accept       json
// @Param        id        path      string           true  "ID de la resource"
// @Param        resource  body      models.Resource  true  "Données mises à jour"
// @Success      204       {string}  string "Mise à jour réussie"
// @Failure      400       {string}  string "Requête invalide"
// @Router       /resources/{id} [put]
func PutResourceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.Context().Value("resourceID").(string)
	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	var body struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	err = collections.UpdateResource(models.Resource{
		ID:   id,
		Name: body.Name,
		URL:  body.URL,
	})
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
