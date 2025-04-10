package helpers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // driver sqlite
	"github.com/sirupsen/logrus"
)

var DB *sql.DB

// InitDB initialise la connexion à SQLite
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./collections.db") // chemin vers le fichier
	if err != nil {
		logrus.Fatal("Échec d'ouverture de la base de données SQLite :", err)
	}

	// On teste la connexion
	err = DB.Ping()
	if err != nil {
		logrus.Fatal("Impossible de ping la base de données :", err)
	}

	logrus.Info("Connexion à la base SQLite réussie.")
}
