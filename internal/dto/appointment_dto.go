package dto

import (
	"api-agendamentos/internal/entity"
	"time"

	"github.com/go-playground/validator/v10"
)

type CreateAppointmentInput struct {
	EmployeeID string    `json:"employeeId" validate:"required"`
	StartDate  time.Time `json:"startDate" validate:"required"`
	FinishDate time.Time `json:"finishDate" validate:"required"`
}

type ListAppointmentsOutput struct {
	List  []entity.Appointment `json:"list"`
	Total int                  `json:"total"`
}

func (u *CreateAppointmentInput) Validate() error {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
