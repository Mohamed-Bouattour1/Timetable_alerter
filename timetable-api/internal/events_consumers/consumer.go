package events_consumers

import (
	"context"
	"encoding/json"
	"time"
	"timetable-api/internal/helpers"
	"timetable-api/internal/models"
	eventService "timetable-api/internal/services/events"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
)

func EventConsumer() (*jetstream.Consumer, error) {
	js, _ := jetstream.New(helpers.NatsConn) // JetStream context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// Récupérer le stream EVENTS
	stream, err := js.Stream(ctx, "EVENTS")
	if err != nil {
		logrus.Error("Stream EVENTS non trouvé :", err)
		return nil, err
	}
	// Vérifier si le consumer durable existe déjà
	consumer, err := stream.Consumer(ctx, "timetable_consumer")
	if err != nil {
		logrus.Warn("ℹConsumer non trouvé -> création...")

		// Création d'un consumer durable
		consumer, err = stream.CreateConsumer(ctx, jetstream.ConsumerConfig{
			Durable:     "timetable_consumer",
			Name:        "timetable_consumer",
			Description: "Consumer durable pour EVENTS.create",
		})
		if err != nil {
			logrus.Error("Erreur création consumer :", err)
			return nil, err
		}

		logrus.Infof("Consumer timetable_consumer créé")
	} else {
		logrus.Infof("Consumer timetable_consumer récupéré (déjà existant)")
	}
	return &consumer, nil
}

func Consume(consumer jetstream.Consumer) (err error) {
	// Consommer les messages
	cc, err := consumer.Consume(func(msg jetstream.Msg) { // à chaque message reçu
		// Lecture du message
		logrus.Debug("Message reçu depuis NATS")
		logrus.Debug(string(msg.Data()))
		// Décoder le JSON en struct Event
		var event models.Event
		err := json.Unmarshal(msg.Data(), &event)
		if err != nil {
			logrus.Error("Erreur de parsing JSON :", err)
			_ = msg.Ack()
			return
		}
		// Appeler le service pour le traitement de l’Event
		err = eventService.ProcessIncomingEvent(event)
		if err != nil {
			logrus.Errorf("Erreur traitement event UID=%s : %v", event.UID, err)
		}

		// Confirme que le message a été traité
		_ = msg.Ack()
	})
	// Maintenir la consommation active
	<-cc.Closed()
	cc.Stop()

	return err
}
