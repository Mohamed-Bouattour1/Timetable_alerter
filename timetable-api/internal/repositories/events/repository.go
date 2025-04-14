package events

import (
	"encoding/json"
	"time"
	"timetable-api/internal/helpers"
	"timetable-api/internal/models"

	"github.com/gofrs/uuid"
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

func GetAllEvents() ([]models.Event, error) {
	rows, err := helpers.DB.Query("SELECT id, resource_ids, uid, description, name, start, end, location, last_update FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Event

	for rows.Next() {
		var e models.Event
		var resourceIds string
		var idStr string
		var start, end, lastUpdate time.Time

		if err := rows.Scan(&idStr, &resourceIds, &e.UID, &e.Description, &e.Name, &start, &end, &e.Location, &lastUpdate); err != nil {
			return nil, err
		}

		e.ID, _ = uuid.FromString(idStr)
		e.Start = start
		e.End = end
		e.LastUpdate = lastUpdate
		json.Unmarshal([]byte(resourceIds), &e.ResourceIds)

		result = append(result, e)
	}

	return result, nil
}

func GetEventByUID(uid string) (models.Event, error) {
	row := helpers.DB.QueryRow("SELECT id, uid, description, name, start, end, location, last_update, resource_ids FROM events WHERE uid = ?", uid)

	var e models.Event
	var idStr, resourceIdsStr string
	var start, end, lastUpdate time.Time

	err := row.Scan(&idStr, &e.UID, &e.Description, &e.Name, &start, &end, &e.Location, &lastUpdate, &resourceIdsStr)
	if err != nil {
		return models.Event{}, err
	}

	e.ID, _ = uuid.FromString(idStr)
	e.Start = start
	e.End = end
	e.LastUpdate = lastUpdate
	json.Unmarshal([]byte(resourceIdsStr), &e.ResourceIds)

	return e, nil
}

func UpdateEvent(e models.Event) error {
	resourceIdsJson, _ := json.Marshal(e.ResourceIds)

	_, err := helpers.DB.Exec(`
		UPDATE events SET 
			description = ?, name = ?, start = ?, end = ?, location = ?, last_update = ?, resource_ids = ?
		WHERE uid = ?
	`,
		e.Description, e.Name, e.Start, e.End, e.Location, e.LastUpdate, string(resourceIdsJson), e.UID)

	return err
}
