package services

import (
	"cpf-cnpj-api/internal/domain/models"
	"cpf-cnpj-api/internal/domain/repositories"
	"cpf-cnpj-api/internal/domain/validators"
	"fmt"
)

type CpfCnpjService struct {
	repository repositories.CpfCnpjRepository
}

func NewCpfCnpjService(repository repositories.CpfCnpjRepository) *CpfCnpjService {
	return &CpfCnpjService{repository: repository}
}

func (s *CpfCnpjService) CreateCpfCnpj(number string, documentType string) (*models.CpfCnpj, error) {
	if err := validators.ValidateCpfCnpj(number, documentType); err != nil {
		return nil, err
	}

	cpfCnpjObj := &models.CpfCnpj{Number: number, Type: models.CpfCnpjType(documentType)}

	err := s.repository.Save(cpfCnpjObj)
	if err != nil {
		return nil, fmt.Errorf("error storing CPF/CNPJ: %w", err)
	}

	return cpfCnpjObj, nil
}

func (s *CpfCnpjService) GetAll() ([]models.CpfCnpj, error) {
	cpfCnpjs, err := s.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error fetching CPFs/CNPJs: %w", err)
	}
	return cpfCnpjs, nil
}
