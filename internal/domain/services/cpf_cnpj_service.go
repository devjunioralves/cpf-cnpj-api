package services

import (
	"cpf-cnpj-api/internal/domain/models"
	"cpf-cnpj-api/internal/domain/repositories"
	"cpf-cnpj-api/internal/domain/validators"
	"cpf-cnpj-api/internal/mappers"
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

	cpfCnpj := &models.CpfCnpj{Number: number, Type: models.CpfCnpjType(documentType)}

	err := s.repository.Save(cpfCnpj)
	if err != nil {
		return nil, fmt.Errorf("error on store CPF/CNPJ: %w", err)
	}

	return cpfCnpj, nil
}

func (s *CpfCnpjService) GetAll() ([]mappers.CpfCnpjDTO, error) {
	entities, err := s.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error on search CPFs/CNPJs: %w", err)
	}

	var dtos []mappers.CpfCnpjDTO
	for _, entity := range entities {
		dtos = append(dtos, mappers.ToCpfCnpjDTO(&entity))
	}

	return dtos, nil
}
