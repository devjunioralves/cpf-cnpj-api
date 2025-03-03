package repositories

import "cpf-cnpj-api/internal/domain/models"

type CpfCnpjRepository interface {
	Save(cpfCnpj *models.CpfCnpj) error
	FindAll() ([]models.CpfCnpj, error)
	BlockCpfCnpjById(id uint) error
}
