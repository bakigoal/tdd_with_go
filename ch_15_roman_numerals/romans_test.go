package ch_15_roman_numerals

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Num  int
		Want string
	}{
		{1, "I"},
	}

	for _, testCase := range cases {
		t.Run(strconv.Itoa(testCase.Num)+" == "+testCase.Want, func(t *testing.T) {
			assert.Equal(t, testCase.Want, ConvertToRoman(testCase.Num))
		})
	}
}
