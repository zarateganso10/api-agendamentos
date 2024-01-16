package dto

import (
	"api-agendamentos/internal/entity"

	"github.com/go-playground/validator/v10"
)

type CreateCompanyInput struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type ListCompaniesOutput struct {
	List  []entity.Company `json:"list"`
	Total int              `json:"total"`
}

func (u *CreateCompanyInput) Validate() error {
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return err
	}
	return nil
}
