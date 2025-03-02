package repositories

import (
	"cpf-cnpj-api/internal/domain/models"

	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	db *gorm.DB
}

func NewUserRepositoryGorm(db *gorm.DB) *UserRepositoryGorm {
	return &UserRepositoryGorm{db: db}
}

func (r *UserRepositoryGorm) Save(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryGorm) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
