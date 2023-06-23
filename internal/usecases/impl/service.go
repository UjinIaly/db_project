package impl

import (
	"github.com/UjinIaly/db_project
/internal/models"
	"github.com/UjinIaly/db_project
/internal/repositories"
	"github.com/UjinIaly/db_project
/internal/usecases"
)

type ServiceUseCaseImpl struct {
	serviceRepository repositories.ServiceRepository
}

func CreateServiceUseCase(serviceRepository repositories.ServiceRepository) usecases.ServiceUseCase {
	return &ServiceUseCaseImpl{serviceRepository: serviceRepository}
}

func (serviceUseCase *ServiceUseCaseImpl) Clear() (err error) {
	return serviceUseCase.serviceRepository.Clear()
}

func (serviceUseCase *ServiceUseCaseImpl) GetStatus() (status *models.Status, err error) {
	return serviceUseCase.serviceRepository.GetStatus()
}
