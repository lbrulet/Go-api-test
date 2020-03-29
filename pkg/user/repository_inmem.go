package user

import (
	"github.com/jinzhu/gorm"

	"github.com/lbrulet/Go-api-test/pkg/models"
)

//InMemRepository is a database in memory structure
type InMemRepository struct {
	m map[int]*models.User
}

//NewInMemRepository is the constructor of InMemRepository
func NewInMemRepository() *InMemRepository {
	var m = map[int]*models.User{}
	return &InMemRepository{m: m}
}

//Migrate is a function to make change on the database
func (r *InMemRepository) Migrate() {
}

//GetUserByID is a method to get a user by ID
func (r *InMemRepository) GetUserByID(ID int) (*models.User, error) {
	if r.m[ID] == nil {
		return nil, gorm.ErrRecordNotFound
	}
	return r.m[ID], nil
}

//DeleteUserByID is a method to delete a user by ID
func (r *InMemRepository) DeleteUserByID(ID int) error {
	if r.m[ID] == nil {
		return gorm.ErrRecordNotFound
	}
	r.m[ID] = nil
	return nil
}

//UpdateUserByID is a method to update a user by ID
func (r *InMemRepository) UpdateUserByID(user *models.User) error {
	if r.m[user.ID] == nil {
		return gorm.ErrRecordNotFound
	}
	r.m[user.ID] = user
	return nil
}

//InsertUser is a method to insert a new user
func (r *InMemRepository) InsertUser(user *models.User) error {
	r.m[user.ID] = user
	return nil
}

//GetAllUsers is a method to get all users
func (r InMemRepository) GetAllUsers() ([]*models.User, error) {
	var d []*models.User
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}
