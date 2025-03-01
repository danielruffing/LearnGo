package iteration

import "strings"

func Repeat(character string, times int) string {
	var result strings.Builder
	for i := 0; i < times; i++ {
		result.WriteString(character)
	}
	return result.String()
}
