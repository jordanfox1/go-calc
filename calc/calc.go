package calc

import (
	"log"
	"strconv"
	"strings"
)

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

func DetermineNumTypes(input string) string {
	if strings.Contains(input, ".") {
		return "float"
	}

	return "int"
}

func CalculateInt(input string, operation string) int {
	input = strings.ReplaceAll(input, " ", "")

	if operation == "add" {
		result := add(input)
		return result
	}

	if operation == "subtract" {
		result := subtract(input)
		return result
	}

	if operation == "multiply" {
		result := multiply(input)
		return result
	}

	return 2
}

func CalculateFloat(input string, operation string) float64 {
	input = strings.ReplaceAll(input, " ", "")

	if operation == "add" {
		result := addFl(input)
		truncatedRes := truncateFloatResult(result)
		return truncatedRes
	}

	return 2
}

func multiply(inp string) int {
	num1, num2 := convertInputStrToNums(inp, "*")
	result := num1 * num2

	return result
}

func subtract(input string) int {
	num1, num2 := convertInputStrToNums(input, "-")
	result := num1 - num2

	return result
}

func add(input string) int {
	num1, num2 := convertInputStrToNums(input, "+")
	result := num1 + num2

	return result
}

func addFl(input string) float64 {
	num1, num2 := convertInputStrToFloats(input, "+")
	result := num1 + num2

	return result
}

func convertInputStrToNums(input string, operator string) (int, int) {
	nums := strings.Split(input, operator)

	if len(nums) != 2 {
		log.Fatal("Invalid input. There should be exactly two numbers separated by '+'.")
	}

	num1, err1 := strconv.Atoi(nums[0])
	num2, err2 := strconv.Atoi(nums[1])

	if err1 != nil || err2 != nil {
		log.Fatal("Invalid input. Both operands should be valid integers.")
	}

	return num1, num2
}

func convertInputStrToFloats(input string, operator string) (float64, float64) {
	nums := strings.Split(input, operator)

	if len(nums) != 2 {
		log.Fatalf("Invalid input. There should be exactly two numbers separated by %s", operator)
	}

	num1, err1 := strconv.ParseFloat(nums[0], 64)
	num2, err2 := strconv.ParseFloat(nums[1], 64)

	if err1 != nil || err2 != nil {
		log.Fatal("Invalid input. Both operands should be valid integers.")
	}

	return num1, num2
}

func truncateFloatResult(inp float64) float64 {
	// Convert the float to a string to manipulate it
	str := strconv.FormatFloat(inp, 'f', -1, 64)

	str = strings.TrimRight(str, "0")

	str = formatZerosFromStr(str)

	str = strings.TrimSuffix(str, ".")

	// Parse the truncated string back to a float64
	truncatedValue, _ := strconv.ParseFloat(str, 64)

	return truncatedValue
}

func formatZerosFromStr(str string) string {
	zeroCount := 0
	newStr := ""

	for _, char := range str {
		if zeroCount >= 3 {
			newStr = strings.TrimRight(newStr, "0")
			return newStr
		}

		if char == '0' {
			zeroCount++
		} else {
			zeroCount = 0
		}

		newStr = newStr + string(char)
	}

	return newStr
}
