package ch_12_reflection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalk(t *testing.T) {

	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	assert.Equal(t, 1, len(got))
	assert.Equal(t, expected, got[0])
}
