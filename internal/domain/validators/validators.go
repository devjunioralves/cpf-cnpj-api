package validators

import "fmt"

var (
	ErrInvalidCPF  = fmt.Errorf("invalid CPF")
	ErrInvalidCNPJ = fmt.Errorf("invalid CNPJ")
)

func ValidateCpfCnpj(number string, documentType string) error {
	if documentType == "CPF" && !ValidateCPF(number) {
		return ErrInvalidCPF
	} else if documentType == "CNPJ" && !ValidateCNPJ(number) {
		return ErrInvalidCNPJ
	}
	return nil
}
