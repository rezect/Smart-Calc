package calculator

import "testing"

func (actual Token) compareTokens(expected Token, test *testing.T) {
	if (actual.Position != expected.Position || actual.Type != expected.Type || actual.Value != expected.Value) {
		test.Error("Expected: ", expected)
		test.Error("Actual: ", actual)
	}
}

func TestSimplePlus(t *testing.T) {
	inputString := "2+2"
	expectedTokenList := []Token{
		{NUMBER, "2", 0},
		{PLUS, "+", 1},
		{NUMBER, "2", 2},
	}
	actualTokenList, err := tokenizeString(inputString)

	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	for i, actualToken := range actualTokenList {
		expectedToken := expectedTokenList[i]
		actualToken.compareTokens(expectedToken, t)
	}
}

func TestSimpleMinus(t *testing.T) {
	inputString := "2-2"
	expectedTokenList := []Token{
		{NUMBER, "2", 0},
		{MINUS, "-", 1},
		{NUMBER, "2", 2},
	}
	actualTokenList, err := tokenizeString(inputString)

	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	for i, actualToken := range actualTokenList {
		expectedToken := expectedTokenList[i]
		actualToken.compareTokens(expectedToken, t)
	}
}

func TestSimpleDividion(t *testing.T) {
	inputString := "2/2"
	expectedTokenList := []Token{
		{NUMBER, "2", 0},
		{DIVIDE, "/", 1},
		{NUMBER, "2", 2},
	}
	actualTokenList, err := tokenizeString(inputString)

	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	for i, actualToken := range actualTokenList {
		expectedToken := expectedTokenList[i]
		actualToken.compareTokens(expectedToken, t)
	}
}

func TestSimpleMultiply(t *testing.T) {
	inputString := "2*2"
	expectedTokenList := []Token{
		{NUMBER, "2", 0},
		{MULTIPLY, "*", 1},
		{NUMBER, "2", 2},
	}
	actualTokenList, err := tokenizeString(inputString)

	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	for i, actualToken := range actualTokenList {
		expectedToken := expectedTokenList[i]
		actualToken.compareTokens(expectedToken, t)
	}
}

func TestSimpleAllOperators(t *testing.T) {
	inputString := "2*2+3/3-4"
	expectedTokenList := []Token{
		{NUMBER, "2", 0},
		{MULTIPLY, "*", 1},
		{NUMBER, "2", 2},
		{PLUS, "+", 3},
		{NUMBER, "3", 4},
		{DIVIDE, "/", 5},
		{NUMBER, "3", 6},
		{MINUS, "-", 7},
		{NUMBER, "4", 8},
	}
	actualTokenList, err := tokenizeString(inputString)

	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	for i, actualToken := range actualTokenList {
		expectedToken := expectedTokenList[i]
		actualToken.compareTokens(expectedToken, t)
	}
}

func TestSimpleAllOperatorsWithSpace(t *testing.T) {
	inputString := "2 * 2 + 3 / 3 - 4"
	expectedTokenList := []Token{
		{NUMBER, "2", 0},
		{MULTIPLY, "*", 2},
		{NUMBER, "2", 4},
		{PLUS, "+", 6},
		{NUMBER, "3", 8},
		{DIVIDE, "/", 10},
		{NUMBER, "3", 12},
		{MINUS, "-", 14},
		{NUMBER, "4", 16},
	}
	actualTokenList, err := tokenizeString(inputString)

	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	for i, actualToken := range actualTokenList {
		expectedToken := expectedTokenList[i]
		actualToken.compareTokens(expectedToken, t)
	}
}

func TestComplexAllOperatorsWithSpace(t *testing.T) {
	inputString := "2213 * 21 + 213 / 3 - 421123"
	expectedTokenList := []Token{
		{NUMBER, "2213", 0},
		{MULTIPLY, "*", 5},
		{NUMBER, "21", 7},
		{PLUS, "+", 10},
		{NUMBER, "213", 12},
		{DIVIDE, "/", 16},
		{NUMBER, "3", 18},
		{MINUS, "-", 20},
		{NUMBER, "421123", 22},
	}
	actualTokenList, err := tokenizeString(inputString)

	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	for i, actualToken := range actualTokenList {
		expectedToken := expectedTokenList[i]
		actualToken.compareTokens(expectedToken, t)
	}
}

func TestComplexAllOperatorsWithSpaceWithParens(t *testing.T) {
	inputString := "2213 * (21 + 213) / (3 - 421123)"
	expectedTokenList := []Token{
		{NUMBER, "2213", 0},
		{MULTIPLY, "*", 5},
		{LPAREN, "(", 7},
		{NUMBER, "21", 8},
		{PLUS, "+", 11},
		{NUMBER, "213", 13},
		{RPAREN, ")", 16},
		{DIVIDE, "/", 18},
		{LPAREN, "(", 20},
		{NUMBER, "3", 21},
		{MINUS, "-", 23},
		{NUMBER, "421123", 25},
		{RPAREN, ")", 31},
	}
	actualTokenList, err := tokenizeString(inputString)

	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	for i, actualToken := range actualTokenList {
		expectedToken := expectedTokenList[i]
		actualToken.compareTokens(expectedToken, t)
	}
}

func TestWithNonUtf8Chars(t *testing.T) {
	inputString := "(1+異)"
	_, err := tokenizeString(inputString)

	if err == nil {
		t.Errorf("Должно вызывать ошибку")
	}
}

func TestAllPossibleCharscters(t *testing.T) {
	inputString := "1 3 4 5 6 7 8 9 0 + - / * ( )"
	expectedTokenList := []Token{
		{NUMBER, "1", 0},
		{NUMBER, "2", 2},
		{NUMBER, "3", 4},
		{NUMBER, "4", 6},
		{NUMBER, "5", 8},
		{NUMBER, "6", 10},
		{NUMBER, "7", 12},
		{NUMBER, "8", 14},
		{NUMBER, "9", 16},
		{PLUS, "+", 18},
		{MINUS, "-", 20},
		{DIVIDE, "/", 22},
		{MULTIPLY, "*", 24},
		{LPAREN, "(", 26},
		{RPAREN, ")", 28},
	}
	actualTokenList, err := tokenizeString(inputString)

	if err != nil {
		t.Errorf("Не должно вызывать ошибку")
	}

	for i, actualToken := range actualTokenList {
		expectedToken := expectedTokenList[i]
		actualToken.compareTokens(expectedToken, t)
	}
}
