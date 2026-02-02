package cli

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"Smart-Calc/internal/calculator"
)

type Config struct {
	Help       bool
	Verbose    bool
	OutputType string // "float", "int", "auto"
	Precision  int
}

func Cli() {
	conf := parseFlags()

	if conf.Help {
		printHelp()
		return
	}

	args := flag.Args()

	if len(args) > 0 {
		handleArgs(conf)
	} else {
		handleStdin(conf)
	}
}

func parseFlags() Config {
	conf := Config{}

	flag.BoolVar(&conf.Help, "help", false, "Показать справочное сообщение")
	flag.BoolVar(&conf.Help, "h", false, "Показать справочное сообщение")
	flag.BoolVar(&conf.Verbose, "verbose", false, "Более подробный вывод")
	flag.BoolVar(&conf.Verbose, "v", false, "Более подробный вывод")
	flag.StringVar(&conf.OutputType, "output", "auto", "Тип выходного значения")
	flag.IntVar(&conf.Precision, "precision", 2, "Количество знаков после запятой")

	flag.Parse()

	return conf
}

func printHelp() {
	fmt.Print(`SmartCalc - калькулятор математических выражений для командной строки

	ИСПОЛЬЗОВАНИЕ:
		smartcalc [ВЫРАЖЕНИЕ]
		smartcalc [ОПЦИИ]

	АРГУМЕНТЫ:
		ВЫРАЖЕНИЕ    Математическое выражение для вычисления (необязательно)
					Если не указано, читается из STDIN.

	ОПЦИИ:
		-help (-h)		Показать эту справку и выйти
		-verbose (-v)	Подробный вывод
		-output	 	 	Тип выходного значения (float, int, auto)
		-precision	 	Количество знаков после запятой ("-1" - автоматически)

	ПОДДЕРЖИВАЕМЫЕ ОПЕРАЦИИ:
		Арифметика:
			+    Сложение
			-    Вычитание (также унарный минус)
			*    Умножение
			/    Деление
			^    Возведение в степень

		Функции:
			sqrt(x)    Квадратный корень
			sin(x)     Синус (в радианах)
			cos(x)     Косинус (в радианах)
			tan(x)     Тангенс (в радианах)
			log(x)     Натуральный логарифм
			exp(x)     Экспонента

		Константы:
			pi         π (3.141592653589793)
			e          Число Эйлера (2.718281828459045)

	ПРИОРИТЕТ ОПЕРАЦИЙ (от высшего к низшему):
		1. () - скобки
		2. -  - унарный минус
		3. ^  - степень
		4. *, / - умножение и деление
		5. +, - - сложение и вычитание

	ВОЗВРАЩАЕМЫЕ КОДЫ:
		0 - успешное выполнение
		1 - ошибка вычисления или синтаксиса

	СООБЩЕНИЯ ОБ ОШИБКАХ:
		Выводятся в STDERR с описанием проблемы:
			- Неправильный синтаксис
			- Деление на ноль
			- Неизвестный оператор/функция
			- Несбалансированные скобки
			- Неправильные аргументы функций

	СМОТРИТЕ ТАКЖЕ:
		Полное техническое задание: README.md в корне проекта

	АВТОР:
		SmartCalc разработан в рамках учебного проекта на Go
	`)
}

func handleArgs(conf Config) {
	equations := flag.Args()

	for _, equation := range equations {
		result, err := calculator.HandleEquation(equation)
		if err != nil {
			fmt.Printf("[ERROR] Incorrect equation: %s\n", equation)
		}
		switch conf.OutputType {
		case "float", "auto":
			fmt.Println("Результат: ", strconv.FormatFloat(result, 'f', conf.Precision, 64))
		case "int":
			fmt.Println("Результат: ", int(result))
		}
	}
}

func handleStdin(conf Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		equation := scanner.Text()
		result, err := calculator.HandleEquation(equation)
		if err != nil {
			fmt.Printf("[ERROR] Неправильное выражение: %s\n", equation)
		}
		switch conf.OutputType {
		case "float", "auto":
			fmt.Println("Результат: ", strconv.FormatFloat(result, 'f', conf.Precision, 64))
		case "int":
			fmt.Println("Результат: ", int(result))
		}
	}
}
