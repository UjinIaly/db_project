package impl

import (
	"db_project/internal/models"
	"db_project/internal/repositories"
	"db_project/internal/usecases"
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
