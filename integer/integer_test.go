package integers

import (
	"fmt"
	"testing"
)

func TestIntegers(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, sum int, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	}

	t.Run("Sum of 2 integers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4
		assertCorrectMessage(t, sum, expected)
	})

}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
