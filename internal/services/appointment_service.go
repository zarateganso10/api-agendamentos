package services

import (
	"api-agendamentos/internal/dto"
	"api-agendamentos/internal/entity"
	"api-agendamentos/internal/repository"
	"net/http"
)

type AppointmentService struct {
	AppointmentRepository *repository.AppointmentRepository
}

func NewAppointmentService(AppointmentRepository *repository.AppointmentRepository) *AppointmentService {
	return &AppointmentService{
		AppointmentRepository: AppointmentRepository,
	}
}

func (service *AppointmentService) ListAppointements(page int, limit int) (*dto.ListAppointmentsOutput, error) {
	listAppointements, err := service.AppointmentRepository.ListAppointements(page, limit)
	if err != nil {
		return nil, err
	}
	total, err := service.AppointmentRepository.GetTotalAppointmentsRows()
	if err != nil {
		return nil, err
	}

	return &dto.ListAppointmentsOutput{
		List:  listAppointements,
		Total: total,
	}, nil
}

func (service *AppointmentService) GetAppointmentByID(id string) (*entity.Appointment, error) {
	appointment, err := service.AppointmentRepository.GetAppointmentByID(id)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

func (service *AppointmentService) CreateAppointment(appointment dto.CreateAppointmentInput) (*entity.Appointment, *dto.ErrorResponse) {
	appointmentCreated, err := service.AppointmentRepository.CreateAppointment(appointment)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return appointmentCreated, nil
}
