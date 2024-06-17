package integers

import "testing"

func TestAdder(t *testing.T) {

	t.Run("adding 1", func(t *testing.T) {
		sum := Add(3, 2)
		expected := 5

		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})

	t.Run("adding 2", func(t *testing.T) {
		sum := Add(38, 2)
		expected := 40

		if sum != expected {
			t.Errorf("expected %d but got %d", expected, sum)
		}
	})

}
