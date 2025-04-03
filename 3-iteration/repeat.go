package iteration

import "strings"

func Repeat(character string, n int) string {
	var repeated string

	for i := 0; i < n; i++ {
		repeated += character
	}
	return repeated
}

func Repeat_v2(character string, n int) {
	var repeated string

	for i := 0; i < n; i++ {
		repeated += character
	}
}

// Faster version using String builder
func Repeat_v3(character string, n int) string {
	var repeated strings.Builder
	for i := 0; i < n; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
