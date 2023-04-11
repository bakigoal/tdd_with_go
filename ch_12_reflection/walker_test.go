package ch_12_reflection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with 1 string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"struct with 2 string fields",
			struct {
				Name string
				City string
			}{"Chris", "NewYork"},
			[]string{"Chris", "NewYork"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 42},
			[]string{"Chris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})
			assert.Equal(t, test.ExpectedCalls, got)
		})
	}
}
