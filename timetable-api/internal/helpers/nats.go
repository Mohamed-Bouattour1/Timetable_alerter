package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"timetable-api/internal/models"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

var NatsConn *nats.Conn
var jsc nats.JetStreamContext

// initialise la connexion à NATS
func InitNats() {
	var err error

	// connexion au serveur NATS
	NatsConn, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.Fatal("Échec de connexion à NATS :", err)
	}

	// utiliser le JetStream
	jsc, err = NatsConn.JetStream()
	if err != nil {
		logrus.Fatal("Erreur JetStream :", err)
	}

	// créer un stream
	_, err = jsc.AddStream(&nats.StreamConfig{
		Name:     "ALERTS",
		Subjects: []string{"ALERTS.*"},
	})
	if err != nil && err != nats.ErrStreamNameAlreadyInUse {
		logrus.Fatal("Erreur création stream :", err)
	}

	logrus.Info("Connexion NATS + stream ALERTS prêts")

}

func PublishAlert(diff models.AlertDiff) error {
	data, err := json.Marshal(diff)
	if err != nil {
		return err
	}

	subject := "ALERTS.create"
	pubAck, err := jsc.PublishAsync(subject, data)
	if err != nil {
		return err
	}

	select {
	case <-pubAck.Ok():
		logrus.Infof("Alerte envoyé sur %s", subject)
		return nil
	case <-pubAck.Err():
		return errors.New(fmt.Sprintf("Échec envoi alerte"))
	}
}
