package iteration

import "strings"

func Repeat(character string, times int) string {
	var repeatCount = times
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
