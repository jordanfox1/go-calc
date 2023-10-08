package calc

import "strings"

func DetermineOperation(input string) string {
	if strings.Contains(input, "+") {
		return "add"
	}
	if strings.Contains(input, "-") {
		return "subtract"
	}
	if strings.Contains(input, "*") {
		return "multiply"
	}
	if strings.Contains(input, "/") {
		return "divide"
	}
	return "invalid"
}
