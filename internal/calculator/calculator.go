package calculator

import (
	"fmt"
	"math"
	"strconv"
)

var tokenTypePriority = map[string]int{
	// Operators
	"+": 1,
	"-": 1,
	"/": 2,
	"*": 2,

	// Functions
	"sin": 9,
	"cos": 9,
}

func HandleEquation(equationString string) error {
	// Разбиваем задачи на токены
	tokens, err := tokenizeString(equationString)
	if err != nil {
		panic(err)
	}

	// Преобразуем в RPN
	tokensRPN, err := tokensToRPN(tokens)
	if err != nil {
		panic(err)
	}

	// Считаем значение выражения
	result, err := calculateEquation(tokensRPN)
	fmt.Printf("Результат: %v", result)

	return nil
}

// 5*8*cos(2+9)+(7*5+8-9*sin(5*5)+5)
func tokensToRPN(tokens []Token) ([]Token, error) {
	var operators = []Token{}
	var rpn = []Token{}

	for _, token := range tokens {
		if token.Type == Number {
			rpn = append(rpn, token)
		} else if t := token.Type; t == Function || t == Lparen {
			operators = append(operators, token)
		} else if token.Type == Operator {
			if len(operators) == 0 {
				operators = append(operators, token)
			} else {
				lastOperator := operators[len(operators)-1]
				if lastOperator.Type == Lparen || tokenTypePriority[lastOperator.Value] < tokenTypePriority[token.Value] {
					operators = append(operators, token)
				} else {
					for len(operators) > 0 {
						lastOperator = operators[len(operators)-1]
						if tokenTypePriority[lastOperator.Value] < tokenTypePriority[token.Value] {
							break
						}
						rpn = append(rpn, operators[len(operators)-1])
						operators = operators[:(len(operators) - 1)]
					}
					operators = append(operators, token)
				}
			}
		} else if token.Type == Rparen {
			if len(operators) == 0 {
				return nil, fmt.Errorf("Неправильное расположение закрывающей скобки")
			}
			for operators[len(operators)-1].Type != Lparen {
				rpn = append(rpn, operators[len(operators)-1])
				operators = operators[:(len(operators) - 1)]
				if len(operators) == 0 {
					return nil, fmt.Errorf("Неправильное расположение закрывающей скобки")
				}
			}
			operators = operators[:len(operators)-1]
		} else if token.Type == Comma {
			continue
		}
	}
	for len(operators) > 0 {
		rpn = append(rpn, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}
	return rpn, nil
}

func calculateEquation(tokens []Token) (float64, error) {
	var stack = []float64{}

	for _, token := range tokens {
		if token.Type == Number {
			floatNumber, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return 0, err
			}

			stack = append(stack, floatNumber)
		} else if t := token.Type; t == Operator || t == Function {
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
	// Операторы
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
		case "*":
			result = a * b
		case "/":
			result = a / b
		}
		*stack = append(*stack, result)
	case Function:
		switch t.Value {
		case "sin":

			if len(*stack) < 1 {
				return fmt.Errorf("В стеке нет чисел для подсчета синуса")
			}

			a := (*stack)[len(*stack)-1]
			*stack = (*stack)[:(len(*stack) - 1)]

			result := math.Sin(a)

			*stack = append(*stack, result)
		case "cos":

			if len(*stack) < 1 {
				return fmt.Errorf("В стеке нет чисел для подсчета косинуса")
			}

			a := (*stack)[len(*stack)-1]
			*stack = (*stack)[:(len(*stack) - 1)]

			result := math.Cos(a)

			*stack = append(*stack, result)
		}
	}

	return nil
}

func calculateArgument(parameterTokens *[]Token, funcName string) (float64, error) {
	if len(*parameterTokens) == 0 {
		return 0, fmt.Errorf("Получен пустой аргумент функции %s", funcName)
	}

	// Считаем аргумент функции
	result, err := calculateEquation(*parameterTokens)
	if err != nil {
		return 0, err
	}

	return result, nil
}
