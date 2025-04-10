package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Timetable représente un emploi du temps
type Resource struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"` // ex: "M1 Groupe 1"
	URL       string    `json:"url"`  // URL pour récupérer les cours
	CreatedAt time.Time `json:"created_at"`
}

// Alert représente une alerte associée à un emploi du temps
type Alert struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`        // email à notifier
	Resource  uuid.UUID `json:"timetable_id"` // clé étrangère vers Timetable
	When      string    `json:"when"`         // always, room, added, removed
	CreatedAt time.Time `json:"created_at"`
}

func ParseUUID(s string) (uuid.UUID, error) {
	return uuid.FromString(s)
}
