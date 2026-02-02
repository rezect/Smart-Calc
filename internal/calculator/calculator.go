package calculator

import (
	"fmt"
	"math"
	"strconv"
)

var tokenTypePriority = map[string]int{
	"+": 1,
	"-": 1,
	"/": 2,
	"*": 2,
}

var functionArgs = map[string]int{
	"sin": 1,
	"cos": 1,
	"log": 2,
}

func HandleEquation(equationString string) error {
	// Разбиваем задачи на токены
	tokens, err := tokenizeString(equationString)
	if err != nil {
		panic(err)
	}

	// Считаем значение выражения
	result, err := calculateEquation(tokens)
	fmt.Printf("Результат: %v", result)

	return nil
}

//5*8*cos(2+9)+(7*5+8-9*sin(5*5)+5)
func tokensToRPN(tokens []Token) ([]Token, error) {
	var operators = []Token{}
	var rpn = []Token{}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token.Type == Number {
			rpn = append(rpn, token)
		} else if token.Type == Function {
			// Убеждаемся, что следующий токен существует и он - Lparen
			if i >= len(tokens)-1 {
				return nil, fmt.Errorf("После функции должны быть скобки")
			} else if tokens[i+1].Type != Lparen {
				return nil, fmt.Errorf("После функции должны быть скобки")
			}

			// Считываем токены до соотв. правой скобки
			// TODO: более качественная обработка ошибок
			var parameterTokens = []Token{}
			var parameters = []float64{}

			parenCounter := 0
			funcToken := token
			var funcResult float64 = 0

			for j := i + 2; j < len(tokens); j++ {
				token = tokens[j]
				if token.Type == Comma {
					argResult, err := calculateArgument(&parameterTokens, funcToken.Value)
					if err != nil {
						return nil, err
					}

					parameters = append(parameters, argResult)
				} else {
					if token.Type == Lparen {
						parenCounter++
					} else if token.Type == Rparen {
						if parenCounter == 0 {
							argResult, err := calculateArgument(&parameterTokens, funcToken.Value)
							if err != nil {
								return nil, err
							}
							
							parameters = append(parameters, argResult)

							switch funcToken.Value {
							case "sin":
								if len(parameters) != 1 {
									return nil, fmt.Errorf("У синуса может быть только один параметр")
								}
								funcResult = math.Sin(parameters[0])
							case "cos":
								if len(parameters) != 1 {
									return nil, fmt.Errorf("У косинуса может быть только один параметр")
								}
								funcResult = math.Cos(parameters[0])
							}

							i = j
							break
						} else {
							parenCounter--
						}
					}
					parameterTokens = append(parameterTokens, token)
				}
			}

			rpn = append(rpn, Token{Number, strconv.FormatFloat(funcResult, 'f', -1, 64), 0})
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
		} else if token.Type == Lparen {
			operators = append(operators, token)
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
	tokensRPN, err := tokensToRPN(tokens)
	if err != nil {
		return 0, err
	}

	for _, token := range tokensRPN {
		if token.Type == Number {
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
			stack = stack[:(len(stack) - 2)]
			localResult := calculateLocalResult(a, b, token.Value)
			stack = append(stack, localResult)
		}
	}

	if len(stack) != 1 {
		return 0, fmt.Errorf("Конечный стек пустой либо переполнен")
	}

	return stack[0], nil
}

func calculateLocalResult(a, b float64, t string) float64 {
	switch t {
	// Операторы
	case "+":
		return a + b
	case "-":
		return a - b
	case "/":
		return a / b
	case "*":
		return a * b
	}

	return 0
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
