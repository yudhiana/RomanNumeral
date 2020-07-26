package main

import (
	"romans_numeric/usecase"
	"testing"
)

func TestConvertNumeralToRoman(t *testing.T) {
	cases := usecase.NewRoman()
	roman := cases.NumberToRoman(1)
	if roman != "I" {
		t.Errorf("Result Should be I , but %v", roman)
	}
}

func TestConvertRomanToNumeral(t *testing.T) {
	cases := usecase.NewRoman()
	numeral := cases.RomanToNumber("I")
	if numeral != 1 {
		t.Errorf("Result shoyld be number 1, but %v", numeral)
	}
}
