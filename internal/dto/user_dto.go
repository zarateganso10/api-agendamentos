package dto

import (
	"api-agendamentos/internal/entity"
	"errors"

	"github.com/go-playground/validator/v10"
)

type CreateUserInput struct {
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Type      string `json:"type" validate:"required"`
	CompanyID string `json:"companyId"`
	Password  string `json:"password" validate:"required"`
}

type ListUsersOutput struct {
	List  []entity.User `json:"list"`
	Total int           `json:"total"`
}

func (u *CreateUserInput) Validate() error {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return err
	}
	if (u.Type == "employee" || u.Type == "company_owner") && u.CompanyID == "" {
		return errors.New("users of type employee needs companyId to be created")
	}
	if u.Type != "employee" && u.Type != "company_owner" && u.CompanyID != "" {
		return errors.New("only employees and company owners can be linked to a company id")
	}
	return nil
}
