package calculator

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenType int

const (
	Number TokenType = iota
	Operator
	Function
	Lparen
	Rparen
	stringToken
	UnaryMinus
)

var operatorsDict = map[rune]TokenType{
	'+': Operator,
	'-': Operator,
	'/': Operator,
	'*': Operator,
	'^': Operator,
	'(': Lparen,
	')': Rparen,
}

var validFunctions = []string{
	"sin",
	"cos",
	"tan",
	"sqrt",
	"log",
	"exp",
}

var constantsValues = map[string]float64{
	"pi": math.Pi,
	"e":  math.E,
}

func parseFunction(t *Token) bool {
	if slices.Contains(validFunctions, t.Value) {
		t.Type = Function
		return true
	}

	return false
}

func parseConstant(t *Token) bool {
	validConstants := make([]string, 0, len(constantsValues))
	for k := range constantsValues {
		validConstants = append(validConstants, k)
	}

	if slices.Contains(validConstants, t.Value) {
		t.Type = Number
		t.Value = strconv.FormatFloat(constantsValues[t.Value], 'f', -1, 64)
		return true
	}

	return false
}

type Token struct {
	Type     TokenType // Number, Operator, Function, Lparen, Rparen
	Value    string
	Position int
}

func tokenizeString(s string) ([]Token, error) {
	if len(s) != utf8.RuneCountInString(s) {
		return nil, fmt.Errorf("Вы использовали неподдерживаемые символы")
	}

	stringNormalization(&s)

	tokenList := make([]Token, 0)
	var curToken = Token{Number, "", 0}
	var parenCounter int = 0

	var allOperators strings.Builder
	for k := range operatorsDict {
		allOperators.WriteString(string(k))
	}

	for i, ch := range s {
		if strings.ContainsRune(allOperators.String(), ch) || ch == ' ' {
			if curToken.Value != "" {
				err := addToken(&curToken, &tokenList)
				if err != nil {
					return nil, err
				}
			}
			if ch != ' ' {
				switch ch {
				case '(':
					parenCounter++
				case ')':
					parenCounter--
				}
				if parenCounter < 0 {
					return nil, fmt.Errorf("Неправильное расположение скобок")
				}

				if ch == '-' && isUnaryMinus(&tokenList) {
					tokenList = append(tokenList, Token{UnaryMinus, "-", i})
					continue
				}

				tokenList = append(tokenList, Token{operatorsDict[ch], string(ch), i})
			}
		} else if unicode.IsDigit(ch) || ch == '.' {
			if ch == '.' && (curToken.Value == "" || curToken.Value[len(curToken.Value)-1] == '.') {
				return nil, fmt.Errorf("Неправильно расположенный разделитель нецелого числа")
			}
			if curToken.Value == "" {
				curToken.Type = Number
				curToken.Position = i
			}

			curToken.Value += string(ch)
		} else if unicode.IsLetter(ch) {
			if curToken.Value == "" {
				curToken.Position = i
			}

			curToken.Type = stringToken
			curToken.Value += string(ch)
		} else {
			return nil, fmt.Errorf("Встречен неизвестный символ на позиции %v", i)
		}
	}

	if curToken.Value != "" {
		err := addToken(&curToken, &tokenList)
		if err != nil {
			return nil, err
		}
	}

	if parenCounter != 0 {
		return nil, fmt.Errorf("Неправильное расположение скобок")
	}

	return tokenList, nil
}

func stringNormalization(s *string) {
	*s = strings.Trim(*s, "\n\r")
}

func addToken(token *Token, tokenList *[]Token) error {
	if token.Type == stringToken && !(parseFunction(token) || parseConstant(token)) {
		return fmt.Errorf("Строка не является ни функцияей, ни константой")
	}
	*tokenList = append(*tokenList, *token)
	token.Type = Number
	token.Value = ""
	token.Position = 0

	return nil
}

func isUnaryMinus(tokenList *[]Token) bool {
	if len(*tokenList) == 0 {
		return true
	}
	if tt := (*tokenList)[len(*tokenList)-1].Type; tt == Lparen || tt == Operator {
		return true
	}
	return false
}
