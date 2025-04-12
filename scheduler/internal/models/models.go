package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// edt
type Resource struct {
	ID        uuid.UUID `json:"id"`         // Identifiant unique de la resource
	Name      string    `json:"name"`       // Nom de la resource
	URL       string    `json:"url"`        // URL pour récupérer l'edt en ICAL
	CreatedAt time.Time `json:"created_at"` // Date de création
}

// un cours dans edt
type Event struct {
	// ID unique de l’événement
	ID uuid.UUID `json:"id"`

	// IDs des resources associées
	ResourceIds []uuid.UUID `json:"resourceIds"`

	// Identifiant unique externe
	UID string `json:"uid"`

	// Description du cours
	Description string `json:"description"`

	// Nom du cours (affiché)
	Name string `json:"name"`

	// Date et heure de début
	Start time.Time `json:"start"`

	// Date et heure de fin
	End time.Time `json:"end"`

	// Salle
	Location string `json:"location"`

	// Dernière mise à jour
	LastUpdate time.Time `json:"lastUpdate"`
}
