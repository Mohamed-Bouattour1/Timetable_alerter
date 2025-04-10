package helpers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

var DB *sql.DB

// init la connexion bd
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./collections.db")
	if err != nil {
		logrus.Fatal("Échec d'ouverture de la base de données SQLite :", err)
	}

	err = DB.Ping()
	if err != nil {
		logrus.Fatal("Impossible de ping la base de données :", err)
	}

	logrus.Info("Connexion à SQLite réussie.")
}
