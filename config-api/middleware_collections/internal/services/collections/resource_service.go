package collections

import (
	"middleware_collections/internal/models"
	repository "middleware_collections/internal/repositories/collections"
	"time"

	"github.com/gofrs/uuid"
)

// CreateResource
func CreateResource(name, url string) (models.Resource, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return models.Resource{}, err
	}

	r := models.Resource{
		ID:        id,
		Name:      name,
		URL:       url,
		CreatedAt: time.Now(),
	}

	err = repository.InsertResource(r)
	return r, err
}

// GetResources
func GetResources() ([]models.Resource, error) {
	return repository.GetAllResources()
}

// GetAlertByID
func GetResourceByID(id uuid.UUID) (*models.Resource, error) {
	return repository.GetResourceByID(id)
}

// UpdateResource
func UpdateResource(r models.Resource) error {
	return repository.UpdateResource(r)
}

// DeleteResource
func DeleteResource(id uuid.UUID) error {
	return repository.DeleteResource(id)
}
