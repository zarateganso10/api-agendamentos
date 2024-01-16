package repository

import (
	"api-agendamentos/internal/dto"
	"api-agendamentos/internal/entity"

	"github.com/jmoiron/sqlx"
)

type CompanyRepository struct {
	ConnectionPG *sqlx.DB
}

func NewCompanyRepository(connectionPG *sqlx.DB) *CompanyRepository {
	return &CompanyRepository{
		ConnectionPG: connectionPG,
	}
}

func (repository *CompanyRepository) ListCompanies(page int, limit int) ([]entity.Company, error) {
	users := []entity.Company{}
	query := "SELECT * FROM companies LIMIT $1 OFFSET $2"
	err := repository.ConnectionPG.Select(&users, query, limit, page*limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *CompanyRepository) GetTotalCompaniesRows() (int, error) {
	var total int
	err := repository.ConnectionPG.Get(&total, "SELECT COUNT(*) FROM companies;")
	if err != nil {
		return -1, err
	}

	return total, nil
}

func (repository *CompanyRepository) GetCompanyByID(id string) (*entity.Company, error) {
	var company entity.Company
	query := "SELECT * FROM companies WHERE id = $1;"
	err := repository.ConnectionPG.Get(&company, query, id)
	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (repository *CompanyRepository) CreateCompany(company dto.CreateCompanyInput) (*entity.Company, error) {
	var companyCreated entity.Company
	query := `
		INSERT INTO companies (name, email) 
		VALUES ($1, $2)
		RETURNING *
	`
	err := repository.ConnectionPG.QueryRowx(query, company.Name, company.Email).StructScan(&companyCreated)
	if err != nil {
		return nil, err
	}
	return &companyCreated, nil
}
