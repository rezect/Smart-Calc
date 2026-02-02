package calculator

import (
	"fmt"
	"math"
	"strconv"
)

func calculateEquation(tokens []Token) (float64, error) {
	var stack = []float64{}

	for _, token := range tokens {
		if token.Type == Number {
			floatNumber, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return 0, err
			}

			stack = append(stack, floatNumber)
		} else if t := token.Type; t == Operator || t == Function || t == UnaryMinus {
			err := calculateLocalResult(&stack, token)
			if err != nil {
				return 0, err
			}
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("Конечный стек пустой либо переполнен")
	}

	return stack[0], nil
}

func calculateLocalResult(stack *[]float64, t Token) error {
	switch t.Type {
	case Operator:

		if len(*stack) < 2 {
			return fmt.Errorf("Невозможно посчитать значение, стек чисел меньше 2 при применении оператора")
		}

		var result float64 = 0
		a, b := (*stack)[len(*stack)-2], (*stack)[len(*stack)-1]
		*stack = (*stack)[:(len(*stack) - 2)]
		switch t.Value {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "/":
			if b == 0 {
				return fmt.Errorf("Dividion by zero")
			}
			result = a / b
		case "*":
			result = a * b
		case "^":
			result = math.Pow(a, b)
		}
		*stack = append(*stack, result)
	case Function:
		if len(*stack) < 1 {
			return fmt.Errorf("В стеке нет чисел для подсчета функции")
		}

		a := (*stack)[len(*stack)-1]
		var result float64 = 0

		switch t.Value {
		case "sin":
			result = math.Sin(a)
		case "cos":
			result = math.Cos(a)
		case "tan":
			result = math.Tan(a)
		case "sqrt":
			result = math.Sqrt(a)
		case "log":
			result = math.Log10(a)
		case "exp":
			result = math.Exp(a)
		}
		(*stack)[(len(*stack) - 1)] = result
	case UnaryMinus:
		if len(*stack) < 1 {
			return fmt.Errorf("В стеке нет чисел для подсчета функции")
		}

		a := (*stack)[len(*stack)-1]
		(*stack)[(len(*stack) - 1)] = -a
	}

	return nil
}
