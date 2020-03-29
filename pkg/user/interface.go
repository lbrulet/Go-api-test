package user

import "github.com/lbrulet/Go-api-test/pkg/models"

//Repository is an interface to handle the user
type Repository interface {
	Migrate()
	GetUserByID(int) (*models.User, error)
	DeleteUserByID(int) error
	UpdateUserByID(*models.User) error
	InsertUser(*models.User) error
	GetAllUsers() ([]*models.User, error)
}