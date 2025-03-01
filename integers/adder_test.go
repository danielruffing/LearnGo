package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		validateSum(t, sum, expected)
	})
	t.Run("", func(t *testing.T) {
		sum := Add(5, 7)
		expected := 12

		validateSum(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
}

func validateSum(t *testing.T, sum, expected int) {
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
