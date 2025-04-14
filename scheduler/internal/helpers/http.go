package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"scheduler/internal/models"

	"github.com/sirupsen/logrus"
)

// récupère les resources depuis l'API Config
func GetResources(apiURL string) ([]models.Resource, error) {
	resp, err := http.Get(apiURL + "/resources")
	if err != nil {
		logrus.Error("Erreur HTTP vers API Config :", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("API Config a retourné le code : %d", resp.StatusCode)
		return nil, errors.New("mauvaise réponse de l'API Config")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var resources []models.Resource
	if err := json.Unmarshal(body, &resources); err != nil {
		logrus.Error("Erreur de décodage JSON :", err)
		return nil, err
	}

	logrus.Infof("%d resources récupérées", len(resources))
	return resources, nil
}

// DownloadICal télécharge le fichier .ics depuis une URL donnée
func DownloadICal(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		logrus.Errorf("Erreur lors du téléchargement de l'URL .ics : %s", url)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Code de réponse inattendu pour l'URL %s : %d", url, resp.StatusCode)
		return nil, errors.New("échec téléchargement du fichier ical")
	}

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Erreur lecture du contenu .ics : %s", err)
		return nil, err
	}

	logrus.Infof("Fichier .ics téléchargé depuis : %s", url)
	return rawData, nil
}
