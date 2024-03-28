package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("adding 2 + 2", func(t *testing.T) {
		result := Add(2, 2)
		expect := 4
		assertCorrectMessage(t, result, expect)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertCorrectMessage(t testing.TB, result, expect int) {
	t.Helper()
	if result != expect {
		t.Errorf("got %d want %d", result, expect)
	}
}
