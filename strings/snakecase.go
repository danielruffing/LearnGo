package strings

import "strings"

func SnakeCase(input string) string {
	//var output = strings.ToLower(input)
	var output = strings.Replace(input, " ", "_", -1)

	return strings.ToLower(output)
}
