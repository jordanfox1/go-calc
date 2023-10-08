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

// func divide(inp string) float32 {

// 	num1, num2 := convertInputStrToNums(inp, "/")
// 	result := float32(num1) / float32(num2)

// 	return result
// }

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
