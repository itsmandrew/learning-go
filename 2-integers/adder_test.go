package integers

import "testing"

func TestAdd(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		got := Add(2, 2)
		want := 4

		assertCorrectMessage(t, got, want)
	})

	t.Run("simple case 2", func(t *testing.T) {
		got := Add(3, 7)
		want := 10

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
