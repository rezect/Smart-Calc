package calculator

import (
	"math"
	"testing"
)

// ТЕСТ ФУНКЦИЙ ТОКЕНИЗАЦИИ
func TestTokensToRPNSimplePlus(t *testing.T) {
	inputTokens := []Token{
		{Number, "2", 0},
		{Operator, "+", 1},
		{Number, "2", 2},
	}
	expectedTokens := []Token{
		{Number, "2", 0},
		{Number, "2", 2},
		{Operator, "+", 1},
	}
	
	actualTokens, err := tokensToRPN(inputTokens)
	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	for i, actualToken := range actualTokens {
		expectedToken := expectedTokens[i]
		actualToken.compareTokens(expectedToken, t)
	}
}

func TestTokensToRPNComplex(t *testing.T) {
	inputTokens := []Token{
		{Number, "5", 0},
		{Operator, "*", 1},
		{Number, "6", 2},
		{Operator, "+", 3},
		{Lparen, "(", 4},
		{Number, "2", 5},
		{Operator, "-", 6},
		{Number, "9", 7},
		{Rparen, ")", 8},
	}
	expectedTokens := []Token{
		{Number, "5", 0},
		{Number, "6", 2},
		{Operator, "*", 1},
		{Number, "2", 5},
		{Number, "9", 7},
		{Operator, "-", 6},
		{Operator, "+", 3},
	}
	
	actualTokens, err := tokensToRPN(inputTokens)
	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	for i, actualToken := range actualTokens {
		expectedToken := expectedTokens[i]
		actualToken.compareTokens(expectedToken, t)
	}
}

// ТЕСТ ФУНКЦИЙ ПОДСЧЕТА RPN
func TestСalculateEquationComplex(t *testing.T) {
	inputStr := "sin(sin((1 + 2) * 0.1) + 0.1)"
	tokens, err := tokenizeString(inputStr)
	if err != nil {
		panic(err)
	}

	// Преобразуем в RPN
	tokensRPN, err := tokensToRPN(tokens)
	if err != nil {
		panic(err)
	}

	// Считаем значение выражения
	actual, err := calculateEquation(tokensRPN)
	var expected float64 = math.Sin(math.Sin(0.3) + 0.1)

	if actual != expected {
		t.Errorf("Должно совпадать: %v", actual)
	}
}

func TestСalculateEquationComplex2(t *testing.T) {
	equationString := "5*8*(2+9)+(7*5+8-9*(5*5)+5)"
	var expected float64 = 263

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
	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	if result != expected {
		t.Errorf("Должно совпадать: %v", result)
	}
}