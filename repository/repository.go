package repository

import (
	"romans_numeric/model"
	"strconv"
)

// Repository godoc
type Repository struct {
}

// NewRepository godoc
func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) dictNumber() *model.Roman {
	return &model.Roman{
		I: 1,
		V: 5,
		X: 10,
		L: 50,
		C: 100,
		D: 500,
		M: 1000,
	}
}

// ToNum godoc
func (r *Repository) ToNum(msg string) int {
	num, _ := strconv.Atoi(msg)
	return num
}

// GetValue godoc
func (r *Repository) GetValue(msg string) int64 {
	var result int64
	switch msg {
	case "I":
		result = r.dictNumber().I
	case "V":
		result = r.dictNumber().V
	case "X":
		result = r.dictNumber().X
	case "L":
		result = r.dictNumber().L
	case "C":
		result = r.dictNumber().C
	case "D":
		result = r.dictNumber().D
	case "M":
		result = r.dictNumber().M
	default:
		result = 0
	}
	return result
}

// GetNumeral godoc
func (r *Repository) GetNumeral() []model.RomanStr {
	dict := []model.RomanStr{
		{Number: 1000, Roman: "M"},
		{Number: 900, Roman: "CM"},
		{Number: 500, Roman: "D"},
		{Number: 400, Roman: "CD"},
		{Number: 100, Roman: "C"},
		{Number: 90, Roman: "XC"},
		{Number: 50, Roman: "L"},
		{Number: 40, Roman: "XL"},
		{Number: 10, Roman: "X"},
		{Number: 9, Roman: "IX"},
		{Number: 5, Roman: "V"},
		{Number: 4, Roman: "IV"},
		{Number: 1, Roman: "I"},
	}
	return dict
}
