package strings

import "testing"

func TestSnakeCase(t *testing.T) {
	t.Run("snake case", func(t *testing.T) {
		want := "snake_case"
		got := SnakeCase("snake case")

		if want != got {
			t.Errorf("want %q but got %q", want, got)
		}
	})

	t.Run("this should be snake case too", func(t *testing.T) {
		want := "this_should_be_snake_case_too"
		got := SnakeCase("this should be snake case too")

		if want != got {
			t.Errorf("want %q but got %q", want, got)
		}
	})

	t.Run("Snake case", func(t *testing.T) {
		want := "snake_case"
		got := SnakeCase("Snake case")

		if want != got {
			t.Errorf("want %q but got %q", want, got)
		}
	})
}
