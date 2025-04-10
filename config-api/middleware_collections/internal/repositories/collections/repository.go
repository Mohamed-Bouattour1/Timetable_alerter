package collections

import (
	"middleware_collections/internal/helpers"
	"middleware_collections/internal/models"
	"time"

	"github.com/gofrs/uuid"
)

// CreateTables crée les tables si elles n'existent pas
func CreateTables() error {
	db := helpers.DB

	resourceTable := `
	CREATE TABLE IF NOT EXISTS resources (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		url TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	alertTable := `
	CREATE TABLE IF NOT EXISTS alerts (
		id TEXT PRIMARY KEY,
		email TEXT NOT NULL,
		resource TEXT NOT NULL,
		"when" TEXT DEFAULT 'always',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(resource) REFERENCES resources(id)
	);
	`

	if _, err := db.Exec(resourceTable); err != nil {
		return err
	}

	if _, err := db.Exec(alertTable); err != nil {
		return err
	}

	return nil
}

// InsertResource ajoute une nouvelle resource
func InsertResource(r models.Resource) error {
	_, err := helpers.DB.Exec(
		"INSERT INTO resources (id, name, url, created_at) VALUES (?, ?, ?, ?)",
		r.ID.String(), r.Name, r.URL, r.CreatedAt,
	)
	return err
}

// InsertAlert ajoute une nouvelle alerte
func InsertAlert(a models.Alert) error {
	_, err := helpers.DB.Exec(
		`INSERT INTO alerts (id, email, resource, "when", created_at) VALUES (?, ?, ?, ?, ?)`,
		a.ID.String(), a.Email, a.Resource.String(), a.When, a.CreatedAt,
	)
	return err
}

// GetAllResources récupère toutes les resources
func GetAllResources() ([]models.Resource, error) {
	rows, err := helpers.DB.Query("SELECT id, name, url, created_at FROM resources")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var resources []models.Resource
	for rows.Next() {
		var r models.Resource
		var created time.Time
		var id string

		if err := rows.Scan(&id, &r.Name, &r.URL, &created); err != nil {
			return nil, err
		}

		r.ID, _ = models.ParseUUID(id)
		r.CreatedAt = created
		resources = append(resources, r)
	}
	return resources, nil
}

// GetAllAlerts récupère toutes les alertes
func GetAllAlerts() ([]models.Alert, error) {
	rows, err := helpers.DB.Query(`SELECT id, email, resource, "when", created_at FROM alerts`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var alerts []models.Alert
	for rows.Next() {
		var a models.Alert
		var created time.Time
		var id, resID string

		if err := rows.Scan(&id, &a.Email, &resID, &a.When, &created); err != nil {
			return nil, err
		}

		a.ID, _ = models.ParseUUID(id)
		a.Resource, _ = models.ParseUUID(resID)
		a.CreatedAt = created
		alerts = append(alerts, a)
	}
	return alerts, nil
}

func UpdateResource(r models.Resource) error {
	_, err := helpers.DB.Exec(
		"UPDATE resources SET name = ?, url = ? WHERE id = ?",
		r.Name, r.URL, r.ID.String(),
	)
	return err
}

func DeleteResource(id uuid.UUID) error {
	_, err := helpers.DB.Exec("DELETE FROM resources WHERE id = ?", id.String())
	return err
}

func UpdateAlert(a models.Alert) error {
	_, err := helpers.DB.Exec(
		`UPDATE alerts SET email = ?, resource = ?, "when" = ? WHERE id = ?`,
		a.Email, a.Resource.String(), a.When, a.ID.String(),
	)
	return err
}

func DeleteAlert(id uuid.UUID) error {
	_, err := helpers.DB.Exec("DELETE FROM alerts WHERE id = ?", id.String())
	return err
}

func GetResourceByID(id uuid.UUID) (*models.Resource, error) {
	row := helpers.DB.QueryRow("SELECT id, name, url, created_at FROM resources WHERE id = ?", id.String())

	var r models.Resource
	var created time.Time
	var idStr string

	err := row.Scan(&idStr, &r.Name, &r.URL, &created)
	if err != nil {
		return nil, err
	}

	r.ID, _ = models.ParseUUID(idStr)
	r.CreatedAt = created
	return &r, nil
}

func GetAlertByID(id uuid.UUID) (*models.Alert, error) {
	row := helpers.DB.QueryRow(`SELECT id, email, resource, "when", created_at FROM alerts WHERE id = ?`, id.String())

	var a models.Alert
	var idStr, resourceStr string
	var created time.Time

	err := row.Scan(&idStr, &a.Email, &resourceStr, &a.When, &created)
	if err != nil {
		return nil, err
	}

	a.ID, _ = models.ParseUUID(idStr)
	a.Resource, _ = models.ParseUUID(resourceStr)
	a.CreatedAt = created
	return &a, nil
}
