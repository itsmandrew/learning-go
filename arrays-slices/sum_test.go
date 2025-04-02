package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		// The syntax for a slice just omits the size when declaring
		numbers := []int{1, 2, 3}

		got := SumDynamic(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// Go does not let u use equality operators with slices.
	// We could write a function to iterate over each got and want slice and check their values
	// We aren't going to do that and just use 'reflect.DeepEqual'

	// Keep in mind also reflect IS NOT type safe, meaning I can compare w/ a string

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll_v2(t *testing.T) {
	got := SumAll_v2([]int{1, 2, 3, 4}, []int{1, 2, 3}, []int{10, 20, 30})
	want := []int{10, 6, 60}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum_Example(t *testing.T) {
	// Arrays have a fixed cap
	// numbers := [5]int{1, 2, 3, 4, 5}

	numbers := [...]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// Keep in mind if I wanted to pass [4]int intto a function that expects [5]int
// it won't compile. Different types
func TestSum_v2(t *testing.T) {

	numbers := [...]int{1, 2, 3, 4, 5}
	got := Sum_v2(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAllTails(t *testing.T) {

	// Helper function, assigning function to a variable
	// Pretty useful when you want to bind functions in 'scope'
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("normal case", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 5}, []int{4, 10})
		want := []int{5, 5, 10}

		checkSums(t, got, want)
	})

	t.Run("edge case where length == 0", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3})
		want := []int{0, 3}

		checkSums(t, got, want)
	})
}
