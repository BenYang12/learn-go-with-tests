package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	expected := "aaaaa"
	repeated := Repeat("a")

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

const repeatCount = 5

func Repeat(character string) string {
	var repeated strings.Builder

	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
