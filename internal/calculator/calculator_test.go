package calculator

import "testing"

// ТЕСТ ФУНКЦИЙ ТОКЕНИЗАЦИИ
func TestTokensToRPNSimplePlus(t *testing.T) {
	inputTokens := []Token{
		{NUMBER, "2", 0},
		{PLUS, "+", 1},
		{NUMBER, "2", 2},
	}
	expectedTokens := []Token{
		{NUMBER, "2", 0},
		{NUMBER, "2", 2},
		{PLUS, "+", 1},
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
		{NUMBER, "5", 0},
		{MULTIPLY, "*", 1},
		{NUMBER, "6", 2},
		{PLUS, "+", 3},
		{LPAREN, "(", 4},
		{NUMBER, "2", 5},
		{MINUS, "-", 6},
		{NUMBER, "9", 7},
		{RPAREN, ")", 8},
	}
	expectedTokens := []Token{
		{NUMBER, "5", 0},
		{NUMBER, "6", 2},
		{MULTIPLY, "*", 1},
		{NUMBER, "2", 5},
		{NUMBER, "9", 7},
		{MINUS, "-", 6},
		{PLUS, "+", 3},
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
	inputTokens := []Token{
		{NUMBER, "5", 0},
		{NUMBER, "6", 2},
		{MULTIPLY, "*", 1},
		{NUMBER, "2", 5},
		{NUMBER, "9", 7},
		{MINUS, "-", 6},
		{PLUS, "+", 3},
	}
	var expected float64 = 23
	
	actual, err := calculateEquation(inputTokens)
	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

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

	// Преобразуем в ОПН
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