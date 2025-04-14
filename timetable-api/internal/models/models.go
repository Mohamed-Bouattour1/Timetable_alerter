package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// un cours dans un emploi du temps
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

type AlertDiff struct {
	Resource uuid.UUID     `json:"resource"`
	Type     string        `json:"type"` // "new" ou "updated"
	Changes  []FieldChange `json:"changes"`
}

type FieldChange struct {
	Field    string `json:"field"`
	OldValue string `json:"old"`
	NewValue string `json:"new"`
}
