package validators

import (
	"strconv"
	"strings"
)

func ValidateCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return false
	}

	for i := 0; i < 10; i++ {
		if strings.Repeat(strconv.Itoa(i), 11) == cpf {
			return false
		}
	}

	return true
}

func ValidateCNPJ(cnpj string) bool {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")

	return len(cnpj) == 14
}
