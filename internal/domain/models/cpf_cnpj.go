package models

import (
	"time"
)

type CpfCnpj struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Number    string      `json:"number" gorm:"unique"`
	Type      CpfCnpjType `json:"type"`
	Blocked   bool        `json:"blocked"`
	CreatedAt time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
}

type CpfCnpjType string

const (
	CpfType  CpfCnpjType = "CPF"
	CnpjType CpfCnpjType = "CNPJ"
)
