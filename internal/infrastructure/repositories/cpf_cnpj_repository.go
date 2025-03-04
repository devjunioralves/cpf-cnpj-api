package repositories

import (
	"cpf-cnpj-api/internal/domain/models"
	"cpf-cnpj-api/internal/infrastructure"
	"fmt"
	"os"

	"gorm.io/gorm"
)

type CpfCnpjRepositoryGorm struct {
	DB *gorm.DB
}

func NewCpfCnpjRepositoryGorm(db *gorm.DB) *CpfCnpjRepositoryGorm {
	return &CpfCnpjRepositoryGorm{DB: db}
}

func (r *CpfCnpjRepositoryGorm) Save(cpfCnpj *models.CpfCnpj) error {
	cryptoKey := []byte(os.Getenv("CRYPTO_KEY"))

	encryptedCpf, err := infrastructure.EncryptData(cpfCnpj.Number, cryptoKey)
	if err != nil {
		return fmt.Errorf("failed to encrypt CPF/CNPJ: %v", err)
	}

	cpfCnpj.Number = encryptedCpf

	return r.DB.Create(cpfCnpj).Error
}

func (r *CpfCnpjRepositoryGorm) FindAll() ([]models.CpfCnpj, error) {
	var cpfs []models.CpfCnpj

	err := r.DB.Find(&cpfs).Error
	if err != nil {
		return nil, err
	}

	cryptoKey := []byte(os.Getenv("CRYPTO_KEY"))

	for i, cpfCnpj := range cpfs {
		decryptedCpf, err := infrastructure.DecryptData(cpfCnpj.Number, cryptoKey)
		if err != nil {
			return nil, fmt.Errorf("failed to decrypt CPF/CNPJ at index %d: %v", i, err)
		}

		cpfs[i].Number = decryptedCpf
	}

	return cpfs, nil
}

func (r *CpfCnpjRepositoryGorm) BlockCpfCnpjById(id uint) error {
	return r.DB.Model(&models.CpfCnpj{}).
		Where("id = ?", id).
		Update("blocked", true).Error
}
