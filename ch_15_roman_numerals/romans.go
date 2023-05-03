package ch_15_roman_numerals

import "strings"

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbols ...byte) uint16 {
	for _, s := range r {
		if s.Symbol == string(symbols) {
			return s.Value
		}
	}
	return 0
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	total := uint16(0)

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]
		// look ahead to next symbol if we can and, the current symbol is base 10 (only valid subtractors)
		notAtEnd := i+1 < len(roman)
		if notAtEnd && isSubtractive(symbol) {
			// get the value of the 2 character string
			if value := allRomanNumerals.ValueOf(symbol, roman[i+1]); value != 0 {
				total += value
				i++
			} else {
				total += allRomanNumerals.ValueOf(symbol)
			}
		} else {
			total += allRomanNumerals.ValueOf(symbol)
		}
	}

	return total
}

func isSubtractive(currentSymbol uint8) bool {
	return currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C'
}
