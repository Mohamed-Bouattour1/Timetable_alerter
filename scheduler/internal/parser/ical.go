package parser

import (
	"bufio"
	"bytes"
	"strings"
	"time"

	"scheduler/internal/models"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// convertit un fichier .ics en une liste d'événements
func ParseICalToEvents(raw []byte, resourceID uuid.UUID) ([]models.Event, error) {
	scanner := bufio.NewScanner(bytes.NewReader(raw))

	var events []models.Event
	currentEvent := make(map[string]string)
	inEvent := false
	var previousLine string

	for scanner.Scan() {
		line := scanner.Text()

		// continuation de la ligne précédente
		if strings.HasPrefix(line, " ") {
			previousLine += strings.TrimLeft(line, " ")
			continue
		}

		// Si ce n’est pas une continuation, on traite la ligne précédente
		if previousLine != "" {
			processICalLine(previousLine, currentEvent)
			previousLine = ""
		}

		// Détection des blocs VEVENT
		if !inEvent && line != "BEGIN:VEVENT" {
			continue
		}

		if line == "BEGIN:VEVENT" {
			inEvent = true
			currentEvent = make(map[string]string)
			continue
		}

		if line == "END:VEVENT" {
			inEvent = false
			if previousLine != "" {
				processICalLine(previousLine, currentEvent)
				previousLine = ""
			}
			event, err := transformToEvent(currentEvent, resourceID)
			if err == nil {
				events = append(events, event)
			}
			continue
		}

		// Stocker la ligne courante pour la traiter ou concaténer
		previousLine = line
	}

	logrus.Infof("%d événements extraits du fichier .ics", len(events))
	return events, nil
}

// Traiter une ligne
func processICalLine(line string, eventMap map[string]string) {
	splitted := strings.SplitN(line, ":", 2)
	if len(splitted) != 2 {
		return
	}
	key := splitted[0]
	value := splitted[1]
	eventMap[key] = value
}

func transformToEvent(data map[string]string, resourceID uuid.UUID) (models.Event, error) {
	id, _ := uuid.NewV4()

	// Lecture des champs iCal
	start, err := time.Parse("20060102T150405Z", data["DTSTART"])
	if err != nil {
		return models.Event{}, err
	}

	end, err := time.Parse("20060102T150405Z", data["DTEND"])
	if err != nil {
		return models.Event{}, err
	}

	lastUpdate, err := time.Parse("20060102T150405Z", data["LAST-MODIFIED"])
	if err != nil {
		lastUpdate = time.Now()
	}

	event := models.Event{
		ID:          id,
		ResourceIds: []uuid.UUID{resourceID},
		UID:         data["UID"],
		Description: data["DESCRIPTION"],
		Name:        data["SUMMARY"],
		Start:       start,
		End:         end,
		Location:    data["LOCATION"],
		LastUpdate:  lastUpdate,
	}

	return event, nil
}
