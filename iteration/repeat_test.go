package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("expect %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {

	b.Run("v1 benchmark for looping", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat("a", 100)
		}
	})

	b.Run("v2 benchmark for looping", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat_v2("a", 100)
		}
	})

	b.Run("v3 benchmark for looping", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Repeat_v3("a", 100)
		}
	})
}
