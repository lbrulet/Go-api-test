package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/lbrulet/Go-api-test/pkg/models"
)

//PostgresRepository is a database structure
type PostgresRepository struct {
	db *gorm.DB
}

//NewPostgresRepository is the constructor of PostgresRepository
func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

//Migrate is a method to make change on the database
func (r *PostgresRepository) Migrate() {
	r.db.AutoMigrate(models.User{})
}

//GetUserByID is a method to get a user by ID
func (r *PostgresRepository) GetUserByID(ID int) (*models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", ID).First(&user).Error
	return &user, err

}

//DeleteUserByID is a method to delete a user by ID
func (r *PostgresRepository) DeleteUserByID(ID int) error {
	return r.db.Where("id = ?", ID).Delete(&models.User{}).Error
}

//UpdateUserByID is a method to update a user by ID
func (r *PostgresRepository) UpdateUserByID(user *models.User) error {
	return r.db.Save(user).Error
}

//InsertUser is a method to insert a user
func (r *PostgresRepository) InsertUser(user *models.User) error {
	return r.db.Create(user).Error
}

//GetAllUsers is a method to get all users
func (r *PostgresRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	return users, err
}