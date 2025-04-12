package events

import (
	"encoding/json"
	"net/http"
	eventService "timetable-api/internal/services/events"

	"github.com/gofrs/uuid"
)

// GetEventByIDHandler godoc
// @Summary      Récupère un événement par ID
// @Tags         events
// @Produce      json
// @Param        id   path      string  true  "ID de l'événement"
// @Success      200  {object}  models.Event
// @Failure      400  {string}  string "ID invalide"
// @Failure      404  {string}  string "Événement non trouvé"
// @Router       /events/{id} [get]
func GetEventByIDHandler(w http.ResponseWriter, r *http.Request) {
	//lire l'ID depuis le context
	idStr := r.Context().Value("eventID").(string)

	// convertir en UUID
	id, err := uuid.FromString(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// appeler le service
	event, err := eventService.GetEventByID(id)
	if err != nil {
		http.Error(w, "Événement non trouvé", http.StatusNotFound)
		return
	}

	// repondre en JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(event)
}
