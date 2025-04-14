package natsclient

import (
	"encoding/json"
	"errors"
	"fmt"

	"scheduler/internal/models"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

var jsc nats.JetStreamContext
var nc *nats.Conn

// init Nats et le stream
func InitStream() {
	var err error

	// connexion au serveur NATS
	nc, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.Fatal("Échec connexion NATS :", err)
	}

	// utiliser le JetStream
	jsc, err = nc.JetStream()
	if err != nil {
		logrus.Fatal("Erreur JetStream :", err)
	}

	// créer un stream
	_, err = jsc.AddStream(&nats.StreamConfig{
		Name:     "EVENTS",
		Subjects: []string{"EVENTS.*"},
	})
	if err != nil && err != nats.ErrStreamNameAlreadyInUse {
		logrus.Fatal("Erreur création stream :", err)
	}

	logrus.Info("Connexion NATS + stream EVENTS prêts")
}

// publier un message
func PublishEvent(event models.Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	subject := "EVENTS.create"
	pubAckFuture, err := jsc.PublishAsync(subject, data)
	if err != nil {
		return err
	}

	select {
	case <-pubAckFuture.Ok():
		logrus.Infof("Event %s envoyé sur %s", event.ID, subject)
		return nil
	case <-pubAckFuture.Err():
		return errors.New(fmt.Sprintf("Échec envoi event %s", event.ID))
	}
}
