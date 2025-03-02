package validators

import (
	"strings"
)

func removeNonDigits(cpf string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, cpf)
}

func calculateDigit(cpf string, weights []int) int {
	var sum int
	for i, weight := range weights {
		sum += int(cpf[i]-'0') * weight
	}
	digit := 11 - (sum % 11)
	if digit >= 10 {
		return 0
	}
	return digit
}

func ValidateCPF(cpf string) bool {
	cpf = removeNonDigits(cpf)

	if len(cpf) != 11 {
		return false
	}

	weights1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}
	d1 := calculateDigit(cpf[:9], weights1)

	if int(cpf[9]-'0') != d1 {
		return false
	}

	weights2 := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}
	d2 := calculateDigit(cpf[:10], weights2)

	return int(cpf[10]-'0') == d2
}
