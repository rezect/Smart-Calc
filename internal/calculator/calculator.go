package calculator

import (
	"fmt"
	"strconv"
)

var tokenTypeValue = map[TokenType]int{
	PLUS:     1,
	MINUS:    1,
	DIVIDE:   2,
	MULTIPLY: 2,
	LPAREN:   0,
	RPAREN:   0,
}

func HandleEquation(equationString string) error {
	// Разбиваем задачи на токены
	tokens, err := tokenizeString(equationString)
	if err != nil {
		panic(err)
	}

	// Преобразуем в ОПН
	tokensRPN, err := tokensToRPN(tokens)
	if err != nil {
		panic(err)
	}

	// Считаем значение выражения
	result, err := calculateEquation(tokensRPN)
	fmt.Printf("Результат: %v", result)

	return nil
}

func tokensToRPN(tokens []Token) ([]Token, error) {
	var operators = []Token{}
	var rpn = []Token{}

	for _, token := range tokens {
		switch token.Type {
		case NUMBER:
			rpn = append(rpn, token)
		default:
			if len(operators) == 0 {
				operators = append(operators, token)
			} else {
				lastOperatorType := operators[len(operators)-1].Type
				if lastOperatorType == LPAREN || token.Type == LPAREN {
					operators = append(operators, token)
				} else if token.Type == RPAREN {
					for operators[len(operators)-1].Type != LPAREN {
						rpn = append(rpn, operators[len(operators)-1])
						operators = operators[:(len(operators) - 1)]
					}
					operators = operators[:len(operators)-1]
				} else if tokenTypeValue[lastOperatorType] < tokenTypeValue[token.Type] {
					operators = append(operators, token)
				} else {
					for len(operators) > 0 {
						lastOperatorType = operators[len(operators)-1].Type
						if tokenTypeValue[lastOperatorType] < tokenTypeValue[token.Type] {
							break
						}
						rpn = append(rpn, operators[len(operators)-1])
						operators = operators[:(len(operators) - 1)]
					}
					operators = append(operators, token)
				}
			}
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
		if token.Type == NUMBER {
			floatNumber, err := strconv.ParseFloat(token.Value, 64)
			if err != nil {
				return 0, err
			}

			stack = append(stack, floatNumber)
		} else {
			if len(stack) < 2 {
				return 0, fmt.Errorf("Невозможно посчитать значение, стек чисел меньше 2 при применении оператора")
			}
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:(len(stack)-2)]
			localResult := calculateLocalResult(a, b, token.Type)
			stack = append(stack, localResult)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("Конечный стек пустой либо переполнен")
	}

	return stack[0], nil
}

func calculateLocalResult(a, b float64, t TokenType) float64 {
	var result float64 = 0
	switch t {
	case PLUS:
		result = a + b
	case MINUS:
		result = a - b
	case MULTIPLY:
		result = a * b
	case DIVIDE:
		result = a / b
	}
	return result
}
