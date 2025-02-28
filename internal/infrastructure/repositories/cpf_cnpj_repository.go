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

func (r *CpfCnpjRepositoryGorm) FindById(id uint) (*models.CpfCnpj, error) {
	var cpf models.CpfCnpj
	err := r.DB.First(&cpf, id).Error
	return &cpf, err
}

func (r *CpfCnpjRepositoryGorm) Update(cpfCnpj *models.CpfCnpj) error {
	return r.DB.Save(cpfCnpj).Error
}

func (r *CpfCnpjRepositoryGorm) Delete(id uint) error {
	return r.DB.Delete(&models.CpfCnpj{}, id).Error
}
