package mappers

import "cpf-cnpj-api/internal/domain/models"

type CpfCnpjDTO struct {
	ID      uint   `json:"id"`
	Number  string `json:"number"`
	Type    string `json:"type"`
	Blocked bool   `json:"blocked"`
}

func ToCpfCnpjDTO(entity *models.CpfCnpj) CpfCnpjDTO {
	return CpfCnpjDTO{
		ID:      entity.ID,
		Number:  entity.Number,
		Type:    string(entity.Type),
		Blocked: entity.Blocked,
	}
}
