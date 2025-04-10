package alert

import (
	"encoding/json"
	alertService "middleware_collections/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// POST /alerts
func PostAlertHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email    string `json:"email"`
		Resource string `json:"resource"`
		When     string `json:"when"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logrus.Error("❌ JSON invalide :", err) //test
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}
	logrus.Infof("📥 Body reçu : %+v", body) //test

	resourceID, err := uuid.FromString(body.Resource)
	if err != nil {
		logrus.Error("❌ UUID ressource invalide :", err) //test
		http.Error(w, "ID de ressource invalide", http.StatusBadRequest)
		return
	}

	alert, err := alertService.CreateAlert(body.Email, resourceID, body.When)
	if err != nil {
		logrus.Error("❌ Erreur lors de la création de l'alerte :", err) //test
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	logrus.Infof("✅ Alerte créée : %+v", alert) // test

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(alert)
}
