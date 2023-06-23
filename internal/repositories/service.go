package repositories

import "db_project/internal/models"

type ServiceRepository interface {
	Clear() (err error)
	GetStatus() (status *models.Status, err error)
}
