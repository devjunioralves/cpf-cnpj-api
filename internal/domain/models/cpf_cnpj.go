package models

import (
	"cpf-cnpj-api/internal/domain/validators"
	"time"
)

type CpfCnpj struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Number    string    `json:"number" gorm:"unique"`
	Type      string    `json:"type"`
	Blocked   bool      `json:"blocked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *CpfCnpj) Validate() bool {
	if c.Type == "CPF" {
		return validators.ValidateCPF(c.Number)
	} else if c.Type == "CNPJ" {
		return validators.ValidateCNPJ(c.Number)
	}
	return false
}
