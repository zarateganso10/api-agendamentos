package services

import (
	"api-agendamentos/configs"
	"api-agendamentos/internal/dto"
	"api-agendamentos/internal/entity"
	repositories "api-agendamentos/internal/repository"
	"api-agendamentos/utils/bcrypt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	UserRepository *repositories.UserRepository
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) CreateSession(auth *entity.Auth) (string, *dto.ErrorResponse) {
	globalConfig := configs.NewParsedConfig()
	expirationTime := time.Now().Add(5 * time.Minute)

	userDB, err := service.UserRepository.GetUserByEmailAndType(auth.Email, auth.Type)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return "", &dto.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "user not found",
			}
		}
		return "", &dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	if !bcrypt.CheckPasswordHash(auth.Password, userDB.Password) {
		return "", &dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "password wrong",
		}
	}

	claims := &Claims{
		Email: auth.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(globalConfig.SecretJWT))
	if err != nil {
		return "", &dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	return tokenString, nil
}
