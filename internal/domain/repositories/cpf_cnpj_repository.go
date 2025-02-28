package repositories

import "cpf-cnpj-api/internal/domain/models"

type CpfCnpjRepository interface {
	Save(cpfCnpj *models.CpfCnpj) error
	FindAll() ([]models.CpfCnpj, error)
	FindById(id uint) (*models.CpfCnpj, error)
	Update(cpfCnpj *models.CpfCnpj) error
	Delete(id uint) error
}
