package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lbrulet/Go-api-test/pkg/models"
)


type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Migrate() {
	r.db.AutoMigrate(models.User{})
}

func (r *PostgresRepository) GetUserByID(ID int) (*models.User, error) {
	var user *models.User
	err := r.db.Where(&models.User{ID: ID}).First(user).Error
	return user, err
}

func (r *PostgresRepository) DeleteUserByID(ID int) error {
	return r.db.Delete(models.User{}, "id =", ID).Error
}

func (r *PostgresRepository) UpdateUserByID(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *PostgresRepository) InsertUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *PostgresRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	return users, err
}