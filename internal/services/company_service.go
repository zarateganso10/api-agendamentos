package services

import (
	"api-agendamentos/internal/dto"
	"api-agendamentos/internal/entity"
	"api-agendamentos/internal/repository"
	"net/http"
)

type CompanyService struct {
	CompanyRepository *repository.CompanyRepository
}

func NewCompanyService(companyRepository *repository.CompanyRepository) *CompanyService {
	return &CompanyService{
		CompanyRepository: companyRepository,
	}
}

func (service *CompanyService) ListCompanies(page int, limit int) (*dto.ListCompaniesOutput, error) {
	listUsers, err := service.CompanyRepository.ListCompanies(page, limit)
	if err != nil {
		return nil, err
	}
	total, err := service.CompanyRepository.GetTotalCompaniesRows()
	if err != nil {
		return nil, err
	}

	return &dto.ListCompaniesOutput{
		List:  listUsers,
		Total: total,
	}, nil
}

func (service *CompanyService) GetCompanyByID(id string) (*entity.Company, error) {
	user, err := service.CompanyRepository.GetCompanyByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *CompanyService) CreateCompany(company dto.CreateCompanyInput) (*entity.Company, *dto.ErrorResponse) {
	companyCreated, err := service.CompanyRepository.CreateCompany(company)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return companyCreated, nil
}
