package alert

import (
	"encoding/json"
	alertService "middleware_collections/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetAlertHandler godoc
// @Summary      Récupère une alerte
// @Tags         alerts
// @Produce      json
// @Param        id   path      string  true  "ID de l'alerte"
// @Success      200  {object}  models.Alert
// @Failure      404  {string}  string "Alerte non trouvée"
// @Router       /alerts/{id} [get]
func GetAlertHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.Context().Value("alertID").(string)
	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}
	result, err := alertService.GetAlertByID(id)
	logrus.Info(result)
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
