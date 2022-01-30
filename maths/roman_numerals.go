package maths

import (
	"strings"
)

type RomanNumeral int

func (rn RomanNumeral) ToInt() int {
	return int(rn)
}

func (rn RomanNumeral) String() string {
	i := int(rn)
	mod := 1000
	ten := true
	var r []string
	ordered := []string{"M", "D", "C", "L", "X", "V", "I"}
	for idx, str := range ordered {
		if ten {
			r = append(r, strings.Repeat(str, i/mod))
			i = i % mod
			if i != 0 && i >= 9*mod/10 {
				r = append(r, ordered[idx+2], str)
				i -= (9 * mod / 10)
			}
			mod = mod / 2
		} else {
			r = append(r, strings.Repeat(str, i/mod))
			i = i % mod
			if i >= 4*mod/5 {
				r = append(r, ordered[idx+1], str)
				i -= (4 * mod / 5)
			}
			mod = mod / 5
		}
		ten = !ten
	}
	return strings.Join(r, "")
}

var (
	numeralMap = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	numeralReplacements = map[string]string{
		"IV": "IIII",
		"IX": "IIIIIIIII",
		"XL": "XXXX",
		"XC": "XXXXXXXXX",
		"CD": "CCCC",
		"CM": "CCCCCCCCC",
	}
)

func NumeralFromString(s string) RomanNumeral {
	for k, v := range numeralReplacements {
		s = strings.ReplaceAll(s, k, v)
	}
	sum := 0
	for i := range s {
		sum += numeralMap[s[i:i+1]]
	}
	return RomanNumeral(sum)
}
