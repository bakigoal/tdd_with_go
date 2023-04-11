package ch_12_reflection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

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
		{
			"nested fields",
			Person{"Chris", Profile{42, "NewYork"}},
			[]string{"Chris", "NewYork"},
		},
		{
			"pointer to things",
			&Person{"Chris", Profile{42, "NewYork"}},
			[]string{"Chris", "NewYork"},
		},
		{
			"slices",
			[]Profile{
				{33, "London"},
				{34, "NewYork"},
			},
			[]string{"London", "NewYork"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
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
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
