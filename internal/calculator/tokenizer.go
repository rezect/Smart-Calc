package calculator

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type TokenType int

const (
	NUMBER TokenType = iota
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	LPAREN
	RPAREN
	EOF
)

var tokenDict = map[rune]TokenType{
	'+': PLUS,
	'-': MINUS,
	'/': DIVIDE,
	'*': MULTIPLY,
	'(': LPAREN,
	')': RPAREN,
}

type Token struct {
	Type     TokenType
	Value    string
	Position int
}

func tokenizeString(s string) ([]Token, error) {
	if len(s) != utf8.RuneCountInString(s) {
		return nil, fmt.Errorf("Вы использовали неподдерживаемые символы")
	}
	
	stringNormalization(&s)

	tokenList := make([]Token, 0)
	var curNumber = Token{NUMBER, "", 0}
	
	for i, ch := range s {
		switch ch {
		case '+', '-', '/', '*', '(', ')', ' ':
			if (curNumber.Value != "") {
				tokenList = append(tokenList, curNumber)
				curNumber.Value = ""
				curNumber.Position = 0
			}
			if (ch != ' ') {
				tokenList = append(tokenList, Token{tokenDict[ch], string(ch), i})
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if curNumber.Value == "" {
				curNumber.Position = i
			}
			curNumber.Value += string(ch)
		default:
			return nil, fmt.Errorf("Встречен неизвестный символ на позиции %v", i)
		}
	}

	if curNumber.Value != "" {
		tokenList = append(tokenList, curNumber)
		curNumber.Value = ""
		curNumber.Position = 0
	}

	return tokenList, nil
}

func stringNormalization(s *string) {
	*s = strings.Trim(*s, "\n\r")
}
