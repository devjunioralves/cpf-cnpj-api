package repositories

import (
	"cpf-cnpj-api/internal/domain/models"

	"gorm.io/gorm"
)

type CpfCnpjRepositoryGorm struct {
	DB *gorm.DB
}

func NewCpfCnpjRepositoryGorm(db *gorm.DB) *CpfCnpjRepositoryGorm {
	return &CpfCnpjRepositoryGorm{DB: db}
}

func (r *CpfCnpjRepositoryGorm) Save(cpfCnpj *models.CpfCnpj) error {
	return r.DB.Create(cpfCnpj).Error
}

func (r *CpfCnpjRepositoryGorm) FindAll() ([]models.CpfCnpj, error) {
	var cpfs []models.CpfCnpj
	err := r.DB.Find(&cpfs).Error
	return cpfs, err
}

func (r *CpfCnpjRepositoryGorm) BlockCpfCnpjById(id uint) error {
	return r.DB.Model(&models.CpfCnpj{}).
		Where("id = ?", id).
		Update("blocked", true).Error
}
