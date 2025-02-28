package validators

import (
	"strings"
)

func ValidateCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return false
	}

	if cpf == "00000000000" || cpf == "11111111111" || cpf == "22222222222" || cpf == "33333333333" || cpf == "44444444444" || cpf == "55555555555" || cpf == "66666666666" || cpf == "77777777777" || cpf == "88888888888" || cpf == "99999999999" {
		return false
	}

	d1 := calculateCpfDigit(cpf[:9])

	d2 := calculateCpfDigit(cpf[:10])

	return cpf[9] == d1 && cpf[10] == d2
}

func calculateCpfDigit(base string) byte {
	var sum int
	weights := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < len(base); i++ {
		sum += int(base[i]-'0') * weights[i]
	}
	rest := sum % 11
	if rest < 2 {
		return '0'
	}
	return byte('0' + 11 - rest)
}
