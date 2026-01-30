package main

import (
	"os"
	"bufio"

	"Smart-Calc/internal/calculator"
)

func main() {
	if (len(os.Args) > 1) {
		err := handleArgs()
		if err != nil {
			panic(err)
		}
	} else {
		err := handleStdin()
		if err != nil {
			panic(err)
		}
	}
}

func handleArgs() error {
	equations := os.Args[1:]
	
	for _, equation := range equations {
		err := calculator.HandleEquation(equation)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func handleStdin() error {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		equation := scanner.Text()
		err := calculator.HandleEquation(equation)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
