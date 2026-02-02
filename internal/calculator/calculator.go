package calculator

import (
	"fmt"
)

func HandleEquation(equationString string) (float64, error) {
	if equationString == "" {
		return 0, fmt.Errorf("Пустая строка")
	}

	tokens, err := tokenizeString(equationString)
	if err != nil {
		return 0, err
	}

	tokensRPN, err := tokensToRPN(tokens)
	if err != nil {
		return 0, err
	}

	result, err := calculateEquation(tokensRPN)
	if err != nil {
		return 0, err
	}

	return result, nil
}
