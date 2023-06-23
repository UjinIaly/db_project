package usecases

import "db_project/internal/models"

type ServiceUseCase interface {
	Clear() (err error)
	GetStatus() (status *models.Status, err error)
}
