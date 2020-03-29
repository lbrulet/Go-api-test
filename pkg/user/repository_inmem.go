package user

import (
	"github.com/jinzhu/gorm"

	"github.com/lbrulet/Go-api-test/pkg/models"
)

type InMemRepository struct {
	m map[int]*models.User
}

func NewInMemRepository() *InMemRepository {
	var m = map[int]*models.User{}
	return &InMemRepository{m: m}
}

func (r *InMemRepository) Migrate() {
}

func (r *InMemRepository) GetUserByID(ID int) (*models.User, error) {
	if r.m[ID] == nil {
		return nil, gorm.ErrRecordNotFound
	}
	return r.m[ID], nil
}

func (r *InMemRepository) DeleteUserByID(ID int) error {
	if r.m[ID] == nil {
		return gorm.ErrRecordNotFound
	}
	r.m[ID] = nil
	return nil
}

func (r *InMemRepository) UpdateUserByID(user *models.User) error {
	if r.m[user.ID] == nil {
		return gorm.ErrRecordNotFound
	}
	r.m[user.ID] = user
	return nil
}

func (r *InMemRepository) InsertUser(user *models.User) error {
	r.m[user.ID] = user
	return nil
}

func (r InMemRepository) GetAllUsers() ([]*models.User, error) {
	var d []*models.User
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}
