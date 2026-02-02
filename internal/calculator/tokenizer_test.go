package calculator

import "testing"

func (actual Token) compareTokens(expected Token, test *testing.T) {
	if (actual.Position != expected.Position || actual.Type != expected.Type || actual.Value != expected.Value) {
		test.Error("Expected: ", expected)
		test.Error("Actual: ", actual)
	}
}

func TestOnlyNumberInt(t *testing.T) {
	inputString := "123"
	expectedTokenList := []Token{
		{Number, "123", 0},
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

func TestOnlyNumberFloat(t *testing.T) {
	inputString := "123.45"
	expectedTokenList := []Token{
		{Number, "123.45", 0},
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

func TestSimplePlus(t *testing.T) {
	inputString := "2+2"
	expectedTokenList := []Token{
		{Number, "2", 0},
		{Operator, "+", 1},
		{Number, "2", 2},
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
		{Number, "2", 0},
		{Operator, "-", 1},
		{Number, "2", 2},
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
		{Number, "2", 0},
		{Operator, "/", 1},
		{Number, "2", 2},
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
		{Number, "2", 0},
		{Operator, "*", 1},
		{Number, "2", 2},
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
		{Number, "2", 0},
		{Operator, "*", 1},
		{Number, "2", 2},
		{Operator, "+", 3},
		{Number, "3", 4},
		{Operator, "/", 5},
		{Number, "3", 6},
		{Operator, "-", 7},
		{Number, "4", 8},
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
		{Number, "2", 0},
		{Operator, "*", 2},
		{Number, "2", 4},
		{Operator, "+", 6},
		{Number, "3", 8},
		{Operator, "/", 10},
		{Number, "3", 12},
		{Operator, "-", 14},
		{Number, "4", 16},
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
		{Number, "2213", 0},
		{Operator, "*", 5},
		{Number, "21", 7},
		{Operator, "+", 10},
		{Number, "213", 12},
		{Operator, "/", 16},
		{Number, "3", 18},
		{Operator, "-", 20},
		{Number, "421123", 22},
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
	inputString := "sin(sin((1 + 2) * 0.1) + 0.1)"
	expectedTokenList := []Token{
		{Function, "sin", 0},
		{Lparen, "(", 3},
		{Function, "sin", 4},
		{Lparen, "(", 7},
		{Lparen, "(", 8},
		{Number, "1", 9},
		{Operator, "+", 11},
		{Number, "2", 13},
		{Rparen, ")", 14},
		{Operator, "*", 16},
		{Number, "0.1", 18},
		{Rparen, ")", 21},
		{Operator, "+", 23},
		{Number, "0.1", 25},
		{Rparen, ")", 28},
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

func TestComplex(t *testing.T) {
	inputString := "5*8*(2+9)+(7*5+8-9*(5*5)+5)"
	expectedTokenList := []Token{
		{Number, "5", 0},
		{Operator, "*", 1},
		{Number, "8", 2},
		{Operator, "*", 3},
		{Lparen, "(", 4},
		{Number, "2", 5},
		{Operator, "+", 6},
		{Number, "9", 7},
		{Rparen, ")", 8},
		{Operator, "+", 9},
		{Lparen, "(", 10},
		{Number, "7", 11},
		{Operator, "*", 12},
		{Number, "5", 13},
		{Operator, "+", 14},
		{Number, "8", 15},
		{Operator, "-", 16},
		{Number, "9", 17},
		{Operator, "*", 18},
		{Lparen, "(", 19},
		{Number, "5", 20},
		{Operator, "*", 21},
		{Number, "5", 22},
		{Rparen, ")", 23},
		{Operator, "+", 24},
		{Number, "5", 25},
		{Rparen, ")", 26},
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
	inputString := "1 2 3 4 5 6 7 8 9 0 + - / * ( )"
	expectedTokenList := []Token{
		{Number, "1", 0},
		{Number, "2", 2},
		{Number, "3", 4},
		{Number, "4", 6},
		{Number, "5", 8},
		{Number, "6", 10},
		{Number, "7", 12},
		{Number, "8", 14},
		{Number, "9", 16},
		{Number, "0", 18},
		{Operator, "+", 20},
		{Operator, "-", 22},
		{Operator, "/", 24},
		{Operator, "*", 26},
		{Lparen, "(", 28},
		{Rparen, ")", 30},
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
