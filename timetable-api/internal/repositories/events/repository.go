package events

import (
	"encoding/json"
	"timetable-api/internal/helpers"
	"timetable-api/internal/models"

	"github.com/sirupsen/logrus"
)

func CreateTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS events (
		id TEXT PRIMARY KEY,
		resource_ids TEXT NOT NULL,
		uid TEXT NOT NULL,
		description TEXT,
		name TEXT,
		start DATETIME NOT NULL,
		end DATETIME NOT NULL,
		location TEXT,
		last_update DATETIME
	);
	`

	_, err := helpers.DB.Exec(query)
	if err != nil {
		logrus.Error("Erreur création table events :", err)
		return err
	}

	logrus.Info("Table events créée ou déjà existante")
	return nil
}
func InsertEvent(e models.Event) error {
	resourceIdsJson, err := json.Marshal(e.ResourceIds)
	if err != nil {
		logrus.Error("Erreur encodage resourceIds :", err)
		return err
	}

	_, err = helpers.DB.Exec(`
		INSERT INTO events (
			id, resource_ids, uid, description, name,
			start, end, location, last_update
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		e.ID.String(),
		string(resourceIdsJson),
		e.UID,
		e.Description,
		e.Name,
		e.Start,
		e.End,
		e.Location,
		e.LastUpdate,
	)

	if err != nil {
		logrus.Error("Erreur insertion event :", err)
		return err
	}

	logrus.Infof("Event inséré : %s", e.ID)
	return nil
}
