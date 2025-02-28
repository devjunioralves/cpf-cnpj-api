package validators

import (
	"strings"
)

func ValidateCNPJ(cnpj string) bool {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")

	if len(cnpj) != 14 {
		return false
	}

	if cnpj == "00000000000000" || cnpj == "11111111111111" || cnpj == "22222222222222" || cnpj == "33333333333333" || cnpj == "44444444444444" || cnpj == "55555555555555" || cnpj == "66666666666666" || cnpj == "77777777777777" || cnpj == "88888888888888" || cnpj == "99999999999999" {
		return false
	}

	d1 := calculateCnpjDigit(cnpj[:12])

	d2 := calculateCnpjDigit(cnpj[:13])

	return cnpj[12] == d1 && cnpj[13] == d2
}

func calculateCnpjDigit(base string) byte {
	var sum int
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	for i := 0; i < len(base); i++ {
		sum += int(base[i]-'0') * weights[i]
	}
	rest := sum % 11
	if rest < 2 {
		return '0'
	}
	return byte('0' + 11 - rest)
}
