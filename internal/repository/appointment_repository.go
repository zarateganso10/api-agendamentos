package repository

import (
	"api-agendamentos/internal/dto"
	"api-agendamentos/internal/entity"

	"github.com/jmoiron/sqlx"
)

type AppointmentRepository struct {
	ConnectionPG *sqlx.DB
}

func NewAppointmentRepository(connectionPG *sqlx.DB) *AppointmentRepository {
	return &AppointmentRepository{
		ConnectionPG: connectionPG,
	}
}

func (repository *AppointmentRepository) ListAppointements(page int, limit int) ([]entity.Appointment, error) {
	appointments := []entity.Appointment{}
	query := "SELECT * FROM appointments LIMIT $1 OFFSET $2"
	err := repository.ConnectionPG.Select(&appointments, query, limit, page*limit)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (repository *AppointmentRepository) GetTotalAppointmentsRows() (int, error) {
	var total int
	err := repository.ConnectionPG.Get(&total, "SELECT COUNT(*) FROM appointments;")
	if err != nil {
		return -1, err
	}

	return total, nil
}

func (repository *AppointmentRepository) GetAppointmentByID(id string) (*entity.Appointment, error) {
	var appointment entity.Appointment
	query := "SELECT * FROM appointments WHERE id = $1;"
	err := repository.ConnectionPG.Get(&appointment, query, id)
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}

func (repository *AppointmentRepository) CreateAppointment(appointment dto.CreateAppointmentInput) (*entity.Appointment, error) {
	var appointmentCreated entity.Appointment
	query := `
		INSERT INTO appointments (employee_id, start_date, finish_date) 
		VALUES ($1, $2, $3)
		RETURNING *
	`
	err := repository.ConnectionPG.QueryRowx(query, appointment.EmployeeID, appointment.StartDate, appointment.FinishDate).StructScan(&appointmentCreated)
	if err != nil {
		return nil, err
	}
	return &appointmentCreated, nil
}
