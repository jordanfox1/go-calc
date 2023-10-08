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

func Calculate(input string, operation string) int {
	input = strings.ReplaceAll(input, " ", "")
	if operation == "add" {
		nums := strings.Split(input, "+")

		if len(nums) != 2 {
			log.Fatal("Invalid input. There should be exactly two numbers separated by '+'.")
		}

		num1, err1 := strconv.Atoi(nums[0])
		num2, err2 := strconv.Atoi(nums[1])

		if err1 != nil || err2 != nil {
			log.Fatal("Invalid input. Both operands should be valid integers.")
		}

		result := num1 + num2

		return result
	}
	return 0
}
