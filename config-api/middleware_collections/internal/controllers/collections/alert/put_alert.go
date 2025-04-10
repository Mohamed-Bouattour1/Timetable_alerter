package alert

import (
	"encoding/json"
	"middleware_collections/internal/models"
	alertService "middleware_collections/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
)

// PUT /alerts/{id}
func PutAlertHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.Context().Value("alertID").(string)
	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	var body struct {
		Email    string `json:"email"`
		Resource string `json:"resource"`
		When     string `json:"when"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	resourceUUID, err := uuid.FromString(body.Resource)
	if err != nil {
		http.Error(w, "Resource UUID invalide", http.StatusBadRequest)
		return
	}

	err = alertService.UpdateAlert(models.Alert{
		ID:       id,
		Email:    body.Email,
		Resource: resourceUUID,
		When:     body.When,
	})
	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
