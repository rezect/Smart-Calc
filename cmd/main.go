package main

import (
	"bufio"
	"fmt"
	"os"

	"Smart-Calc/internal/calculator"
)

func main() {
	if (len(os.Args) > 1) {
		handleArgs()
	} else {
		handleStdin()
	}
}

func handleArgs() {
	equations := os.Args[1:]
	
	for _, equation := range equations {
		_, err := calculator.HandleEquation(equation)
		if err != nil {
			fmt.Printf("[ERROR] Incorrect equation: %s\n", equation)
		}
	}
}

func handleStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		equation := scanner.Text()
		_, err := calculator.HandleEquation(equation)
		if err != nil {
			fmt.Printf("[ERROR] Incorrect equation: %s\n", equation)
		}
	}
}
