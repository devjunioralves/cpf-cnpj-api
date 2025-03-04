package validators

import (
	"fmt"
	"strings"
)

var (
	ErrInvalidCPF  = fmt.Errorf("invalid CPF")
	ErrInvalidCNPJ = fmt.Errorf("invalid CNPJ")
)

func ValidateCpfCnpj(number string, documentType string) error {
	documentType = strings.ToUpper(documentType)

	if documentType == "CPF" && !ValidateCPF(number) {
		return ErrInvalidCPF
	} else if documentType == "CNPJ" {
		valid, err := ValidateCNPJ(number)
		if err != nil || !valid {
			return ErrInvalidCNPJ
		}
	}
	return nil
}
