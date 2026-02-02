package calculator

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestHandleEquation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected float64
		wantErr  bool
	}{
		// Базовые операции
		{
			name:     "simple addition",
			input:    "2 + 3",
			expected: 5,
			wantErr:  false,
		},
		{
			name:     "simple subtraction",
			input:    "10 - 4",
			expected: 6,
			wantErr:  false,
		},
		{
			name:     "simple multiplication",
			input:    "6 * 7",
			expected: 42,
			wantErr:  false,
		},
		{
			name:     "simple division",
			input:    "15 / 3",
			expected: 5,
			wantErr:  false,
		},
		{
			name:     "division with decimal result",
			input:    "10 / 4",
			expected: 2.5,
			wantErr:  false,
		},

		// Операции с плавающей точкой
		{
			name:     "float addition",
			input:    "2.5 + 3.7",
			expected: 6.2,
			wantErr:  false,
		},
		{
			name:     "float multiplication",
			input:    "2.5 * 4",
			expected: 10,
			wantErr:  false,
		},

		// Приоритет операций
		{
			name:     "multiplication before addition",
			input:    "2 + 3 * 4",
			expected: 14,
			wantErr:  false,
		},
		{
			name:     "division before subtraction",
			input:    "10 - 6 / 2",
			expected: 7,
			wantErr:  false,
		},
		{
			name:     "parentheses change priority",
			input:    "(2 + 3) * 4",
			expected: 20,
			wantErr:  false,
		},

		// Сложные выражения со скобками
		{
			name:     "nested parentheses",
			input:    "((2 + 3) * (4 - 1)) / 2",
			expected: 7.5,
			wantErr:  false,
		},
		{
			name:     "multiple parentheses levels",
			input:    "(1 + (2 * (3 + 4))) - 5",
			expected: 10,
			wantErr:  false,
		},

		// Математические функции
		{
			name:     "simple sin",
			input:    "sin(0)",
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "simple cos",
			input:    "cos(0)",
			expected: 1,
			wantErr:  false,
		},
		{
			name:     "simple tan",
			input:    "tan(0)",
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "sqrt function",
			input:    "sqrt(16)",
			expected: 4,
			wantErr:  false,
		},
		{
			name:     "log function",
			input:    "log(100)",
			expected: 2,
			wantErr:  false,
		},
		{
			name:     "exp function",
			input:    "exp(0)",
			expected: 1,
			wantErr:  false,
		},

		// Комбинации функций и операций
		{
			name:     "function in expression",
			input:    "sin(0) + cos(0)",
			expected: 1,
			wantErr:  false,
		},
		{
			name:     "function with operation inside",
			input:    "sin(pi/2)",
			expected: 1,
			wantErr:  false,
		},
		{
			name:     "nested functions",
			input:    "sin(cos(0))",
			expected: math.Sin(1),
			wantErr:  false,
		},
		{
			name:     "function with complex argument",
			input:    "sqrt(9 + 16)",
			expected: 5,
			wantErr:  false,
		},

		// Константы
		{
			name:     "pi constant",
			input:    "pi",
			expected: math.Pi,
			wantErr:  false,
		},
		{
			name:     "e constant",
			input:    "e",
			expected: math.E,
			wantErr:  false,
		},
		{
			name:     "expression with pi",
			input:    "pi * 2",
			expected: math.Pi * 2,
			wantErr:  false,
		},

		// // Экспоненциальная запись
		// {
		// 	name:     "scientific notation",
		// 	input:    "1.5e2 + 50",
		// 	expected: 200,
		// 	wantErr:  false,
		// },
		// {
		// 	name:     "negative scientific notation",
		// 	input:    "2e-1 * 10",
		// 	expected: 2,
		// 	wantErr:  false,
		// },

		// Отрицательные числа
		{
			name:     "negative number addition",
			input:    "-5 + 10",
			expected: 5,
			wantErr:  false,
		},
		{
			name:     "negative multiplication",
			input:    "-3 * -4",
			expected: 12,
			wantErr:  false,
		},
		{
			name:     "negative in parentheses",
			input:    "(-2 + 5) * 3",
			expected: 9,
			wantErr:  false,
		},

		// Граничные случаи и особые значения
		{
			name:     "division by one",
			input:    "5 / 1",
			expected: 5,
			wantErr:  false,
		},
		{
			name:     "multiplication by zero",
			input:    "42 * 0",
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "zero addition",
			input:    "0 + 7",
			expected: 7,
			wantErr:  false,
		},

		// Ошибки и некорректный ввод
		{
			name:     "empty string",
			input:    "",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "invalid characters",
			input:    "2 @ 3",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "unmatched parentheses",
			input:    "(2 + 3",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "extra parentheses",
			input:    "2 + 3)",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "division by zero",
			input:    "5 / 0",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "unknown function",
			input:    "unknown(5)",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "invalid function syntax",
			input:    "sin 5",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "multiple operators",
			input:    "2 + + 3",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "operator at start",
			input:    "+ 5",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "operator at end",
			input:    "5 +",
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := HandleEquation(tt.input)

			if tt.wantErr {
				assert.Error(t, err, "Expected error for input: %s", tt.input)
			} else {
				assert.NoError(t, err, "Unexpected error for input: %s", tt.input)
				assert.InDelta(t, tt.expected, actual, 1e-9,
					"For input: %s, expected: %v, got: %v", tt.input, tt.expected, actual)
			}
		})
	}
}
