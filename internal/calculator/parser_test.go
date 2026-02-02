package calculator

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// isNewTokenMoreImportant
func TestIsNewTokenMoreImportant_OperatorsFalse(t *testing.T) {
	actual := isNewTokenMoreImportant(Token{Operator, "+", 0}, Token{Operator, "*", 0})
	expected := false

	assert.Equal(t, actual, expected)
}

func TestIsNewTokenMoreImportant_OperatorsEquals(t *testing.T) {
	actual := isNewTokenMoreImportant(Token{Operator, "+", 0}, Token{Operator, "+", 0})
	expected := false

	assert.Equal(t, actual, expected)
}

func TestIsNewTokenMoreImportant_OperatorsTrue(t *testing.T) {
	actual := isNewTokenMoreImportant(Token{Operator, "/", 0}, Token{Operator, "-", 0})
	expected := true

	assert.Equal(t, actual, expected)
}

func TestIsNewTokenMoreImportant_FunctionOperator(t *testing.T) {
	actual := isNewTokenMoreImportant(Token{Function, "sin", 0}, Token{Operator, "-", 0})
	expected := true

	assert.Equal(t, actual, expected)
}

func TestIsNewTokenMoreImportant_OperatorFunction(t *testing.T) {
	actual := isNewTokenMoreImportant(Token{Operator, "/", 0}, Token{Function, "sin", 0})
	expected := false

	assert.Equal(t, actual, expected)
}

// tokensToRPN
func TestTokensToRPN_SimplePlus(t *testing.T) {
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
	
	assert.Nil(t, err)
	assert.Equal(t, len(actualTokens), len(expectedTokens))

	for i, actualToken := range actualTokens {
		expectedToken := expectedTokens[i]
		actualToken.compareTokens(expectedToken, t)
	}
}

func TestTokensToRPN_Complex(t *testing.T) {
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
		{Operator, "+", 0},
		{Function, "sin", 0},
		{Lparen, "(", 0},
		{Number, "90", 0},
		{Rparen, ")", 0},
	}
	expectedTokens := []Token{
		{Number, "5", 0},
		{Number, "6", 2},
		{Operator, "*", 1},
		{Number, "2", 5},
		{Number, "9", 7},
		{Operator, "-", 6},
		{Operator, "+", 3},
		{Number, "90", 0},
		{Function, "sin", 0},
		{Operator, "+", 0},
	}
	
	actualTokens, err := tokensToRPN(inputTokens)
	
	assert.Nil(t, err)
	assert.Equal(t, len(actualTokens), len(expectedTokens))

	for i, actualToken := range actualTokens {
		expectedToken := expectedTokens[i]
		actualToken.compareTokens(expectedToken, t)
	}
}