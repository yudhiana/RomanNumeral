package usecase

import (
	"romans_numeric/repository"
)

// RomanRepo godoc
type RomanRepo struct {
}

// NewRoman godoc
func NewRoman() *RomanRepo {
	return &RomanRepo{}
}

// RomanToNumber godoc
func (r *RomanRepo) RomanToNumber(msg string) int64 {
	var result int64
	numeral := repository.NewRepository()
	for i := 0; i < len(msg); i++ {
		if i < (len(msg) - 1) {
			if numeral.GetValue(string(msg[i])) < numeral.GetValue(string(msg[i+1])) {
				result += numeral.GetValue(string(msg[i+1])) - numeral.GetValue(string(msg[i]))
				i++
			} else {
				result += numeral.GetValue(string(msg[i]))
			}
		} else {
			result += numeral.GetValue(string(msg[i]))
		}
	}
	return result
}

// NumberToRoman godoc
func (r *RomanRepo) NumberToRoman(msg int) string {
	var result string
	var i int
	romans := repository.NewRepository()
	for msg > 0 {
		for j := 0; j < (msg / romans.GetNumeral()[i].Number); j++ {
			result += romans.GetNumeral()[i].Roman
			msg -= romans.GetNumeral()[i].Number
		}
		i++
	}
	return result
}
