package events

import (
	"fmt"
	"time"
	natsclient "timetable-api/internal/helpers"
	"timetable-api/internal/models"
	"timetable-api/internal/repositories/events"

	"github.com/sirupsen/logrus"
)

func GetEvents() ([]models.Event, error) {
	return events.GetAllEvents()
}

func ProcessIncomingEvent(newEvent models.Event) error {
	// Chercher un event existant avec le même UID
	oldEvent, err := events.GetEventByUID(newEvent.UID)
	if err != nil {
		logrus.Infof("Nouvel événement UID=%s -> insertion + alerte", newEvent.UID)
		err = events.InsertEvent(newEvent) // événement nouveau
		if err != nil {
			return err
		}
		diff := models.AlertDiff{
			Resource: newEvent.ResourceIds[0],
			Type:     "new",
			Changes:  []models.FieldChange{},
		}

		return natsclient.PublishAlert(diff)
	}

	// Il existe -> on le compare
	changed, updatedEvent, diff := CompareEvents(oldEvent, newEvent)
	if !changed {
		logrus.Infof("Événement UID=%s inchangé", newEvent.UID)
		return nil
	}
	for _, change := range diff.Changes {
		logrus.Infof("Champ changé: %s — de [%s] à [%s]", change.Field, change.OldValue, change.NewValue)
	}

	logrus.Warnf("Événement UID=%s modifié -> update + alerte", newEvent.UID)

	// mettre à jour LastUpdate
	updatedEvent.LastUpdate = time.Now()

	err = events.UpdateEvent(updatedEvent)
	if err != nil {
		logrus.Error("Échec update event :", err)
		return err
	}
	return natsclient.PublishAlert(diff)

}

func CompareEvents(old, new models.Event) (bool, models.Event, models.AlertDiff) {
	var changes []models.FieldChange
	modified := false

	// Comparer la localisation
	if old.Location != new.Location {

		changes = append(changes, models.FieldChange{
			Field:    "location",
			OldValue: old.Location,
			NewValue: new.Location,
		})
		modified = true
	}
	// Comparer les resourceIds (cas special)
	for _, r := range new.ResourceIds {
		found := false
		for _, oldR := range old.ResourceIds {
			if r == oldR {
				found = true
				break
			}
		}
		if !found {
			logrus.Infof("Nouvelle resource ajoutée à UID=%s", new.UID)

			// fusionner les resources
			new.ResourceIds = append(old.ResourceIds, r)

			changes = append(changes, models.FieldChange{
				Field:    "resourceIds",
				OldValue: fmt.Sprint(old.ResourceIds),
				NewValue: fmt.Sprint(new.ResourceIds),
			})
			modified = true
		}
	}

	// Comparer start
	if !old.Start.Equal(new.Start) {
		changes = append(changes, models.FieldChange{
			Field:    "start",
			OldValue: old.Start.String(),
			NewValue: new.Start.String(),
		})
		modified = true
	}

	// Comparer end
	if !old.End.Equal(new.End) {
		changes = append(changes, models.FieldChange{
			Field:    "end",
			OldValue: old.End.String(),
			NewValue: new.End.String(),
		})
		modified = true
	}

	if modified {
		diff := models.AlertDiff{
			Resource: new.ResourceIds[0],
			Type:     "updated",
			Changes:  changes,
		}
		return true, new, diff
	}

	return false, new, models.AlertDiff{}
}
