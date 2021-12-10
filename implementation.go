package lab2

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func isOperator(s string) (result bool) {
	var operators = []string{"+", "-", "*", "/", "^"}
	for i := 0; i < len(operators); i++ {
		if s == operators[i] {
			return true
		}
	}
	return false
}

func isNumber(s string) (result bool) {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func stringToNumber(s string) (result float64) {
	result, _ = strconv.ParseFloat(s, 64)
	return result
}

func getOperationResult(first float64, second float64, operation string) (result float64) {
	if operation == "+" {
		result = first + second
	} else if operation == "-" {
		result = first - second
	} else if operation == "*" {
		result = first * second
	} else if operation == "/" {
		result = first / second
	} else if operation == "^" {
		result = math.Pow(first, second)
	}
	return result
}

func PostfixExpressionCalculator(input string) (float64, error) {
	inputSymbols := strings.Split((strings.Trim(input, " ")), " ")
	var stack []float64

	if len(inputSymbols) < 3 {
		return 0, fmt.Errorf("invalid input(short input expression)")
	}

	for i := 0; i < len(inputSymbols); i++ {
		nextSymbol := inputSymbols[i]

		if isNumber(nextSymbol) {
			stack = append(stack, stringToNumber(nextSymbol))
		} else if isOperator(nextSymbol) {
			if len(stack) < 2 {
				return 0, fmt.Errorf("invalid input(too many operators in the expression)")
			}

			secondOperand := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			firstOperand := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			operationResult := getOperationResult(firstOperand, secondOperand, nextSymbol)
			stack = append(stack, operationResult)
			continue
		} else {
			return 0, fmt.Errorf("invalid input expression")
		}
	}

	if len(stack) > 1 {
		return 0, fmt.Errorf("invalid input expression(missed operator(s))")
	}

	result := stack[len(stack)-1]
	return result, nil
}
