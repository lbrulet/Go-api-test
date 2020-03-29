package user

import "github.com/lbrulet/Go-api-test/pkg/models"

//Service is a structure to store the user repository
type Service struct {
	r Repository
}

//NewService is the constructor of Service
func NewService(r Repository) *Service {
	return &Service{r: r}
}

//Migrate is a method to make change on the database
func (s *Service) Migrate() {
	s.r.Migrate()
}

//GetUserByID is a method to get a user by ID
func (s *Service) GetUserByID(ID int) (*models.User, error) {
	return s.r.GetUserByID(ID)
}

//DeleteUserByID is a method to delete a user by ID
func (s *Service) DeleteUserByID(ID int) error {
	return s.r.DeleteUserByID(ID)
}

//UpdateUserByID is a method to update a user by ID
func (s *Service) UpdateUserByID(user *models.User) error {
	return s.r.UpdateUserByID(user)
}

//InsertUser is a method to insert a user
func (s *Service) InsertUser(user *models.User) error {
	return s.r.InsertUser(user)
}

//GetAllUsers is a method to get all users
func (s *Service) GetAllUsers() ([]*models.User, error) {
	return s.r.GetAllUsers()
}