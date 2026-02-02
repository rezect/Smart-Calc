package calculator

import (
	"fmt"
	"strconv"
)

func HandleEquation(equationString string) (float64, error) {
	if equationString == "" {
		return 0, fmt.Errorf("Пустая строка")
	}

	// Разбиваем задачи на токены
	tokens, err := tokenizeString(equationString)
	if err != nil {
		return 0, err
	}

	// Преобразуем в RPN
	tokensRPN, err := tokensToRPN(tokens)
	if err != nil {
		return 0, err
	}

	// Считаем значение выражения
	result, err := calculateEquation(tokensRPN)
	if err != nil {
		return 0, err
	}
	fmt.Println("Результат: ", strconv.FormatFloat(result, 'f', -1, 64))

	return result, nil
}
