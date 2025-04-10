package events

import (
	"encoding/json"
	"net/http"
	"time"
	eventService "timetable-api/internal/services/events"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// PostEventHandler godoc
// @Summary      Crée un événement
// @Description  Ajoute un événement (cours) dans la base
// @Tags         events
// @Accept       json
// @Produce      json
// @Param        event  body      models.Event  true  "Nouvel événement à créer"
// @Success      201    {object}  models.Event
// @Failure      400    {string}  string "Requête invalide"
// @Failure      500    {string}  string "Erreur serveur"
// @Router       /events [post]
func PostEventHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		ResourceIds []string `json:"resourceIds"`
		UID         string   `json:"uid"`
		Description string   `json:"description"`
		Name        string   `json:"name"`
		Start       string   `json:"start"`
		End         string   `json:"end"`
		Location    string   `json:"location"`
		LastUpdate  string   `json:"lastUpdate"`
	}

	// Lire le JSON
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logrus.Error("JSON invalide :", err)
		http.Error(w, "Format JSON invalide", http.StatusBadRequest)
		return
	}

	// Conversion des resourceIds en UUID
	var resourceUUIDs []uuid.UUID
	for _, id := range body.ResourceIds {
		u, err := uuid.FromString(id)
		if err != nil {
			logrus.Error("UUID invalide :", err)
			http.Error(w, "UUID resource invalide", http.StatusBadRequest)
			return
		}
		resourceUUIDs = append(resourceUUIDs, u)
	}

	// Conversion des dates
	start, err := time.Parse(time.RFC3339, body.Start)
	if err != nil {
		http.Error(w, "Date de début invalide", http.StatusBadRequest)
		return
	}

	end, err := time.Parse(time.RFC3339, body.End)
	if err != nil {
		http.Error(w, "Date de fin invalide", http.StatusBadRequest)
		return
	}

	lastUpdate, err := time.Parse(time.RFC3339, body.LastUpdate)
	if err != nil {
		http.Error(w, "Date de mise à jour invalide", http.StatusBadRequest)
		return
	}

	// Appel du service
	event, err := eventService.CreateEvent(
		resourceUUIDs,
		body.UID,
		body.Description,
		body.Name,
		start,
		end,
		body.Location,
		lastUpdate,
	)

	if err != nil {
		http.Error(w, "Erreur serveur", http.StatusInternalServerError)
		return
	}

	// Répondre avec le JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}
