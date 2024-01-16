package repository

import (
	"api-agendamentos/internal/dto"
	"api-agendamentos/internal/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	ConnectionPG *sqlx.DB
}

func NewUserRepository(connectionPG *sqlx.DB) *UserRepository {
	return &UserRepository{
		ConnectionPG: connectionPG,
	}
}

func (repository *UserRepository) ListUsers(page int, limit int) ([]entity.User, error) {
	users := []entity.User{}
	query := "SELECT * FROM users LIMIT $1 OFFSET $2"
	err := repository.ConnectionPG.Select(&users, query, limit, page*limit)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *UserRepository) GetUserByID(id string) (*entity.User, error) {
	var user entity.User
	query := "SELECT * FROM users WHERE id = $1;"
	err := repository.ConnectionPG.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) GetUserByEmailAndType(email string, userType string) (*entity.User, error) {
	var user entity.User
	err := repository.ConnectionPG.Get(
		&user,
		`
			SELECT * FROM users
			WHERE email = $1 AND type = $2;
		`,
		email, userType,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) GetTotalUsersRows() (int, error) {
	var total int
	err := repository.ConnectionPG.Get(&total, "SELECT COUNT(*) FROM users")
	if err != nil {
		return -1, err
	}

	return total, nil
}

func (repository *UserRepository) CreateUser(user dto.CreateUserInput) (*entity.User, error) {
	var userCreated entity.User
	query := `
		INSERT INTO users (name, type, email, password) 
		VALUES ($1, $2, $3, $4)
		RETURNING *
	`

	err := repository.ConnectionPG.QueryRowx(query, user.Name, user.Type, user.Email, user.Password).StructScan(&userCreated)
	if err != nil {
		return nil, err
	}
	return &userCreated, nil
}

func (repository *UserRepository) CreateEmployeeUser(user dto.CreateUserInput) (*entity.User, error) {
	var userCreated entity.User
	query := `
		INSERT INTO users (name, type, email, password, company_id) 
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`

	err := repository.ConnectionPG.QueryRowx(query, user.Name, user.Type, user.Email, user.Password, user.CompanyID).StructScan(&userCreated)
	if err != nil {
		return nil, err
	}
	return &userCreated, nil
}

func (repository *UserRepository) DeleteUserById(id string) error {
	_, err := repository.ConnectionPG.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (repository *UserRepository) VerifyIfUserExistsByEmailAndType(email string, userType string) (bool, error) {
	var userExists bool
	err := repository.ConnectionPG.QueryRow(
		`SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND type = $2)`,
		email, userType,
	).Scan(&userExists)
	if err != nil {
		return false, err
	}
	return userExists, nil
}

func (repository *UserRepository) VerifyIfTypeExists(userType string) (bool, error) {
	var typeExists bool
	err := repository.ConnectionPG.QueryRow(
		`SELECT EXISTS(SELECT 1 FROM users_type WHERE name = $1)`,
		userType,
	).Scan(&typeExists)
	if err != nil {
		return false, err
	}
	return typeExists, nil
}
