package ch_15_roman_numerals

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Arabic    int
		WantRoman string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
	}

	for _, testCase := range cases {
		t.Run(strconv.Itoa(testCase.Arabic)+" == "+testCase.WantRoman, func(t *testing.T) {
			assert.Equal(t, testCase.WantRoman, ConvertToRoman(testCase.Arabic))
		})
	}
}
