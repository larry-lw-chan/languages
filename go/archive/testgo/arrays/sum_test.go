package arrays

import (
	"fmt"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, result, expect int) {
		t.Helper()
		if result != expect {
			t.Errorf("got %d want %d", result, expect)
		}
	}
	t.Run("numbers should all Sum correctly to 15", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sum(numbers)
		expect := 15
		assertCorrectMessage(t, result, expect)
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sum(numbers)
		expect := 6
		assertCorrectMessage(t, result, expect)
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func TestSumAll(t *testing.T) {
	t.Run("add two collection and return Summed array", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func ExampleSumAll() {
	sum := SumAll([]int{1, 2}, []int{0, 9})
	fmt.Println(sum)
	// Output: [3 9]
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("sum the 'tails' of each collection array", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("make sure guards implemented against empty arrays", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	sum := SumAllTails([]int{1, 2}, []int{0, 9})
	fmt.Println(sum)
	// Output: [2 9]
}
