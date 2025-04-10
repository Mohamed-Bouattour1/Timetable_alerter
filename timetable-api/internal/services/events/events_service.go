package events

import (
	"time"
	"timetable-api/internal/models"
	"timetable-api/internal/repositories/events"

	"github.com/gofrs/uuid"
)

// créer un nouvel événement
func CreateEvent(
	resourceIds []uuid.UUID,
	uid string,
	description string,
	name string,
	start time.Time,
	end time.Time,
	location string,
	lastUpdate time.Time,
) (models.Event, error) {
	// Générer un nouvel UUID
	id, err := uuid.NewV4()
	if err != nil {
		return models.Event{}, err
	}

	// Construire l'objet Event
	event := models.Event{
		ID:          id,
		ResourceIds: resourceIds,
		UID:         uid,
		Description: description,
		Name:        name,
		Start:       start,
		End:         end,
		Location:    location,
		LastUpdate:  lastUpdate,
	}

	// Enregistrer dans la base
	err = events.InsertEvent(event)
	return event, err
}
