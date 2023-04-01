package ch_04_arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if got != want {
		t.Errorf("Got '%d' want '%d' given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3}, []int{4, 5, 6})
	want := []int{4, 15}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 3}, []int{4, 5, 6})
		want := []int{3, 11}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5, 6})
		want := []int{0, 11}
		checkSums(t, got, want)
	})
}
