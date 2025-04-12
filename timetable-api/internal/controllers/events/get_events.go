package events

import (
	"encoding/json"
	"net/http"
	eventService "timetable-api/internal/services/events"
)

// GetAllEventsHandler godoc
// @Summary      Liste tous les événements
// @Tags         events
// @Produce      json
// @Success      200  {array}   models.Event
// @Failure      500  {string}  string "Erreur serveur"
// @Router       /events [get]
func GetAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := eventService.GetEvents()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des événements", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}
