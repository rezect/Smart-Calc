package calculator

import (
	"fmt"
)

var tokenTypePriority = map[string]int{
	// Operators
	"+": 1,
	"-": 1,
	"/": 2,
	"*": 2,
	"^": 3,

	// Functions
	"sin":  9,
	"cos":  9,
	"tan":  9,
	"sqrt": 9,
	"log":  9,
	"exp":  9,
}

func isNewTokenMoreImportant(new Token, last Token) bool {
	switch new.Type {
	case Function:
		switch last.Type {
		case Function, Lparen, Rparen:
			return false
		case Operator, UnaryMinus:
			return true
		default:
			panic("Пришло ни функция, ни оператор")
		}
	case Operator:
		switch last.Type {
		case Function, UnaryMinus, Lparen, Rparen:
			return false
		case Operator:
			return tokenTypePriority[new.Value] > tokenTypePriority[last.Value]
		default:
			panic("Пришло ни функция, ни оператор")
		}
	case UnaryMinus:
		switch last.Type {
		case Function, UnaryMinus, Lparen, Rparen:
			return false
		case Operator:
			return true
		default:
			panic("Пришло ни функция, ни оператор")
		}
	default:
		panic("Пришло ни функция, ни оператор")
	}
}

func tokensToRPN(tokens []Token) ([]Token, error) {
	var operators = []Token{}
	var rpn = []Token{}

	var isPrevTokenIsFunction bool = false

	for _, token := range tokens {
		if isPrevTokenIsFunction {
			if token.Type != Lparen {
				return nil, fmt.Errorf("После функции должны открываться скобки")
			} else {
				isPrevTokenIsFunction = false
			}
		}

		if token.Type == Number {
			rpn = append(rpn, token)
		} else if token.Type == Lparen {
			operators = append(operators, token)
		} else if token.Type == Function {
			operators = append(operators, token)
			isPrevTokenIsFunction = true
		} else if token.Type == Operator || token.Type == UnaryMinus {
			if len(operators) == 0 {
				operators = append(operators, token)
			} else {
				lastOperator := operators[len(operators)-1]
				if lastOperator.Type == Lparen || isNewTokenMoreImportant(token, lastOperator) {
					operators = append(operators, token)
				} else {
					for len(operators) > 0 {
						lastOperator = operators[len(operators)-1]
						if lastOperator.Type == Lparen || isNewTokenMoreImportant(token, lastOperator) {
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
		}
	}
	for len(operators) > 0 {
		rpn = append(rpn, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}
	return rpn, nil
}
