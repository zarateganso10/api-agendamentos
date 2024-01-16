package services

import (
	"api-agendamentos/internal/dto"
	"api-agendamentos/internal/entity"
	"api-agendamentos/internal/repository"
	"api-agendamentos/utils/bcrypt"
	"net/http"
)

type UserService struct {
	UserRepository    *repository.UserRepository
	CompanyRepository *repository.CompanyRepository
}

func NewUserService(userRepository *repository.UserRepository, companyRepository *repository.CompanyRepository) *UserService {
	return &UserService{
		UserRepository:    userRepository,
		CompanyRepository: companyRepository,
	}
}

func (service *UserService) ListUsers(page int, limit int) (*dto.ListUsersOutput, error) {
	listUsers, err := service.UserRepository.ListUsers(page, limit)
	if err != nil {
		return nil, err
	}
	total, err := service.UserRepository.GetTotalUsersRows()
	if err != nil {
		return nil, err
	}

	return &dto.ListUsersOutput{
		List:  listUsers,
		Total: total,
	}, nil
}

func (service *UserService) GetUserByID(id string) (*entity.User, error) {
	user, err := service.UserRepository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) CreateUser(user dto.CreateUserInput) (*entity.User, *dto.ErrorResponse) {
	err := user.Validate()
	if err != nil {
		return nil, &dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	typeExists, err := service.UserRepository.VerifyIfTypeExists(user.Type)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if !typeExists {
		return nil, &dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "type of user does not exists",
		}
	}
	if user.CompanyID != "" {
		_, err = service.CompanyRepository.GetCompanyByID(user.CompanyID)
		if err != nil {
			return nil, &dto.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "company id doesn't exists",
			}
		}
	}
	userExists, err := service.UserRepository.VerifyIfUserExistsByEmailAndType(user.Email, user.Type)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if userExists {
		return nil, &dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "user already exists",
		}
	}
	passwordHash, err := bcrypt.HashPassword(user.Password)
	if err != nil {
		return nil, &dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	user.Password = passwordHash
	isNotUserOfTypeEmployee := user.Type != "employee" && user.Type != "company_owner"
	if isNotUserOfTypeEmployee {
		userCreated, err := service.UserRepository.CreateUser(user)
		if err != nil {
			return nil, &dto.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		return userCreated, nil
	} else {
		userCreated, err := service.UserRepository.CreateEmployeeUser(user)
		if err != nil {
			return nil, &dto.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		return userCreated, nil
	}

}

func (service *UserService) UpdateUser(id string, user *entity.User) (*entity.User, error) {
	return nil, nil
}

func (service *UserService) DeleteUserById(id string) error {
	err := service.UserRepository.DeleteUserById(id)
	return err
}
