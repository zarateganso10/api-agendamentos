package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        string        `db:"id" json:"id"`
	Name      string        `db:"name" json:"name"`
	Email     string        `db:"email" json:"email"`
	Type      string        `db:"type" json:"type"`
	Password  string        `db:"password" json:"-"`
	CompanyID uuid.NullUUID `db:"company_id" json:"companyId"`
	CreatedAt time.Time     `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time     `db:"updated_at" json:"updatedAt"`
}
