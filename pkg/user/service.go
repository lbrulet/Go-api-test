package user

import "github.com/lbrulet/Go-api-test/pkg/models"

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r: r}
}

func (s *Service) Migrate() {
	s.r.Migrate()
}

func (s *Service) GetUserByID(ID int) (*models.User, error) {
	return s.r.GetUserByID(ID)
}

func (s *Service) DeleteUserByID(ID int) error {
	return s.r.DeleteUserByID(ID)
}

func (s *Service) UpdateUserByID(user *models.User) error {
	return s.r.UpdateUserByID(user)
}

func (s *Service) InsertUser(user *models.User) error {
	return s.r.InsertUser(user)
}

func (s *Service) GetAllUsers() ([]*models.User, error) {
	return s.r.GetAllUsers()
}




