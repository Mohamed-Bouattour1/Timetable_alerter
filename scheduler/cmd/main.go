package main

import (
	"context"
	"os"
	"os/signal"
	"scheduler/internal/helpers"
	"scheduler/internal/natsclient"
	"scheduler/internal/parser"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/zhashkevych/scheduler"
)

func main() {
	logrus.Info("Scheduler démarré...")

	ctx := context.Background()
	sc := scheduler.NewScheduler()
	sc.Add(ctx, recoverAndSendEvents, time.Second*30)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	sc.Stop()
	logrus.Info("Scheduler arrêté.")
}
func recoverAndSendEvents(ctx context.Context) {
	logrus.Info("Récupération et envoi des événements...")

	// Récupérer toutes les resources depuis l'API Config
	resources, err := helpers.GetResources("http://localhost:8080")
	if err != nil {
		logrus.Error("Échec récupération des resources :", err)
		return
	}

	// Pour chaque resource, télécharger son .ics
	for _, resource := range resources {
		logrus.Infof("Traitement de la resource : %s (%s)", resource.Name, resource.ID)

		rawData, err := helpers.DownloadICal(resource.URL)
		if err != nil {
			logrus.Errorf("Erreur téléchargement .ics pour %s", resource.Name)
			continue
		}

		// Parser le contenu .ics pour obtenir une liste d'événements
		events, err := parser.ParseICalToEvents(rawData, resource.ID)
		if err != nil {
			logrus.Errorf("Erreur parsing .ics pour %s", resource.Name)
			continue
		}

		logrus.Infof("%d événements extraits pour %s", len(events), resource.Name)

		// Publier chaque événement dans NATS
		for _, e := range events {
			err := natsclient.PublishEvent(e)
			if err != nil {
				logrus.Errorf("Échec publication NATS pour event %s", e.ID)
			}
		}
	}
}
