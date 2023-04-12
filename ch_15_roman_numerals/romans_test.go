package ch_15_roman_numerals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	got := ConvertToRoman(1)
	want := "I"
	assert.Equal(t, want, got)
}
