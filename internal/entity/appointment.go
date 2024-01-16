package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Appointment struct {
	ID         string        `db:"id" json:"id"`
	ClientID   uuid.NullUUID `db:"client_id" json:"clientId"`
	EmployeeID string        `db:"employee_id" json:"employeeId"`
	StartDate  time.Time     `db:"start_date" json:"startDate"`
	FinishDate time.Time     `db:"finish_date" json:"finishDate"`
	CreatedAt  time.Time     `db:"created_at" json:"createdAt"`
	UpdatedAt  time.Time     `db:"updated_at" json:"updatedAt"`
}
