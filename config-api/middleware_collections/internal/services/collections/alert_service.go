package collections

import (
	"middleware_collections/internal/models"
	repository "middleware_collections/internal/repositories/collections"
	"time"

	"github.com/gofrs/uuid"
)

// CreateAlert
func CreateAlert(email string, resource uuid.UUID, when string) (models.Alert, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return models.Alert{}, err
	}

	a := models.Alert{
		ID:        id,
		Email:     email,
		Resource:  resource,
		When:      when,
		CreatedAt: time.Now(),
	}

	err = repository.InsertAlert(a)
	return a, err
}

// GetAlerts
func GetAlerts() ([]models.Alert, error) {
	return repository.GetAllAlerts()
}

// GetResourceByID
func GetResourceByID(id uuid.UUID) (*models.Resource, error) {
	return repository.GetResourceByID(id)
}

// UpdateAlert
func UpdateAlert(a models.Alert) error {
	return repository.UpdateAlert(a)
}

// DeleteAlert
func DeleteAlert(id uuid.UUID) error {
	return repository.DeleteAlert(id)
}
