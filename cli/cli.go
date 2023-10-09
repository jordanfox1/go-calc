package cli

import (
	"bufio"
	"fmt"
	"go-calc/calc"
	custom_error "go-calc/error"
	"os"
	"strings"
)

func PromptForInput() error {
	fmt.Println("Welcome to the calculator!")

	fmt.Println("Please enter a sum with two numbers separated by an operator: ")

	// Read user input from the console
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	validInput, err := validateInput(input)
	if err != nil {
		return err
	}

	numType := calc.DetermineNumTypes(validInput)
	operation := calc.DetermineOperation(validInput)

	if numType == "float" {
		result := calc.CalculateFloat(validInput, operation)
		fmt.Println("result: ", result)
		return nil
	}

	result := calc.CalculateInt(validInput, operation)
	fmt.Println("result: ", result)

	return nil
}

func validateInput(input string) (string, error) {
	var numOfOperators uint16

	for _, char := range input {
		switch char {
		case '+', '-', '*', '/':
			numOfOperators++
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', ' ', '.':
			// Numeric character, do nothing
		default:
			// Non-numeric character detected
			err := &custom_error.CustomError{
				Message: "Non-numeric input detected!!",
			}
			fmt.Println(err.Error())
			return "invalid", err
		}
	}

	if numOfOperators != 1 {
		err := &custom_error.CustomError{
			Message: "Invalid number of operators!!",
		}
		fmt.Println(err.Error())
		return "invalid", err
	}

	return input, nil
}
