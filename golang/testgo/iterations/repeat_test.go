package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeating 'a' five times", func(t *testing.T) {
		result := Repeat("a")
		expect := "aaaaa"
		assertCorrectMessage(t, result, expect)
	})
}

func ExampleRepeat() {
	repeated := Repeat("a")
	fmt.Println(repeated)
	// Output: aaaaa
}

func assertCorrectMessage(t testing.TB, result, expect string) {
	t.Helper()
	if result != expect {
		t.Errorf("got %q want %q", result, expect)
	}
}

// Benchmarking
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
