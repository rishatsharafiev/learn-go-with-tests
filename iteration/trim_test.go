package iteration

import (
	"fmt"
	"testing"
	"unicode"
)

func TestTrimFunc(t *testing.T) {
	trimed := TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	expected := "HelloGophers"

	if trimed != expected {
		t.Errorf("expected %q but got %q", expected, trimed)
	}
}

func BenchmarkTrimFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})
	}
}

func ExampleTrimFunc() {
	result := TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	fmt.Println(result)
	// Output: HelloGophers
}
