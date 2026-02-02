package calculator

import (
	"fmt"
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
	Comma
)

var operatorsDict = map[rune]TokenType{
	'+': Operator,
	'-': Operator,
	'/': Operator,
	'*': Operator,
	'(': Lparen,
	')': Rparen,
	',': Comma,
}

var validFunctions = []string{
	"sin",
	"cos",
}

func isFunctionValid(fName string) bool {
	for _, validFuncName := range validFunctions {
		if validFuncName == fName {
			return true
		}
	}
	return false
}

type Token struct {
	Type     TokenType	// Number, Operator, Function, Lparen, Rparen
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

	var allOperators strings.Builder
	for k := range operatorsDict {
		allOperators.WriteString(string(k))
	}
	
	for i, ch := range s {
		if strings.ContainsRune(allOperators.String(), ch) || ch == ' ' {
			if (curToken.Value != "") {
				err := addToken(&curToken, &tokenList)
				if err != nil {
					return nil, err
				}
			}
			if (ch != ' ') {
				tokenList = append(tokenList, Token{operatorsDict[ch], string(ch), i})
			}
		} else if unicode.IsDigit(ch) || ch == '.' {
			if ch == '.' && curToken.Value == "" {
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

			curToken.Type = Function
			curToken.Value += string(ch)
		} else {
			return nil, fmt.Errorf("Встречен неизвестный символ на позиции %v", i)
		}
	}

	if curToken.Value != "" {
		tokenList = append(tokenList, curToken)
	}

	return tokenList, nil
}

func stringNormalization(s *string) {
	*s = strings.Trim(*s, "\n\r")
}

func addToken(token *Token, tokenList *[]Token) error {
	if token.Type == Function {
		if !isFunctionValid(token.Value) {
			return fmt.Errorf("Неизвестное название функции: %v", token.Value)
		}
	}
	*tokenList = append(*tokenList, *token)
	token.Type = Number
	token.Value = ""
	token.Position = 0

	return nil
}
