package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"middleware_collections/internal"
	"middleware_collections/internal/helpers"
	"middleware_collections/internal/repositories/collections"
)

func main() {
	// Init BDD SQLite
	helpers.InitDB()

	// Création des tables si elles n’existent pas
	if err := collections.CreateTables(); err != nil {
		logrus.Fatal("Erreur création des tables :", err)
	}

	// Création des routes
	r := internal.SetupRoutes()

	// Port 8080 si pas de var d'env
	port := os.Getenv("API_CONFIG_PORT")
	if port == "" {
		port = "8080"
		logrus.Warn("Aucune variable d'environnement 'API_CONFIG_PORT', utilisation de 8080 par défaut.")
	}

	logrus.Infof("Serveur lancé sur http://localhost:%s", port)
	logrus.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
