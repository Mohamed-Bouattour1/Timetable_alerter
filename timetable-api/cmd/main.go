package main

import (
	"fmt"
	"net/http"
	"os"
	"timetable-api/internal"
	events_consumers "timetable-api/internal/events_consumers"
	"timetable-api/internal/helpers"
	"timetable-api/internal/repositories/events"

	"github.com/sirupsen/logrus"
)

func main() {
	//Connexion à SQLite
	helpers.InitDB()

	// Création des tables
	if err := events.CreateTables(); err != nil {
		logrus.Fatal("Erreur création des tables :", err)
	}

	// Connexion à NATS
	helpers.InitNats()

	// Lancer le consumer dans une Go routine
	go func() {
		logrus.Info("📡 Lancement du consumer NATS...")
		consumer, err := events_consumers.EventConsumer()
		if err != nil {
			logrus.Warnf("❌ Erreur création consumer : %v", err)
			return
		}
		err = events_consumers.Consume(*consumer)
		if err != nil {
			logrus.Warnf("❌ Erreur consommation NATS : %v", err)
		}
	}()

	//Définir les routes
	r := internal.SetupRoutes()

	//Lire le port
	port := os.Getenv("API_TIMETABLE_PORT")
	if port == "" {
		port = "8080"
		logrus.Warn("Aucun port défini, utilisation du port 8080 par défaut")
	}

	//Démarrer le serveur
	logrus.Infof("Serveur lancé sur http://localhost:%s", port)
	logrus.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
