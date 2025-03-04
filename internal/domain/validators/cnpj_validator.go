package validators

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var invalidCNPJs = map[string]struct{}{
	"00000000000000": {},
	"11111111111111": {},
	"22222222222222": {},
	"33333333333333": {},
	"44444444444444": {},
	"55555555555555": {},
	"66666666666666": {},
	"77777777777777": {},
	"88888888888888": {},
	"99999999999999": {},
}

func ValidateCNPJ(digits string) (bool, error) {
	return valid(digits)
}

func sanitize(data string) string {
	var builder strings.Builder
	for _, r := range data {
		if r >= '0' && r <= '9' {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func valid(data string) (bool, error) {
	data = sanitize(data)

	if len(data) != 14 {
		return false, errors.New("invalid length")
	}

	if _, found := invalidCNPJs[data]; found || !validateCNPJChecksum(data) {
		return false, errors.New("invalid value")
	}

	return true, nil
}

func validateCNPJChecksum(data string) bool {
	digits := stringToIntSlice(data)

	return verifyChecksum(digits, 5, 12) && verifyChecksum(digits, 6, 13)
}

func stringToIntSlice(data string) []int {
	var res []int
	for i := 0; i < len(data); i++ {
		x, _ := strconv.Atoi(string(data[i]))
		res = append(res, x)
	}
	return res
}

func verifyChecksum(data []int, startIdx int, endIdx int) bool {
	soma := 0
	j := startIdx

	for i := 0; i < endIdx; i++ {
		soma += data[i] * j
		if j == 2 {
			j = 9
		} else {
			j--
		}
	}

	resto := soma % 11
	expectedDigit := 0
	if resto >= 2 {
		expectedDigit = 11 - resto
	}

	if data[endIdx] != expectedDigit {
		return false
	}

	return true
}

func FilterNumber(text string) string {
	re := regexp.MustCompile("[0-9]+")
	result := re.FindAllString(text, -1)
	return strings.Join(result, "")
}
