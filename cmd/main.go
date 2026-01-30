package main

import (
	"bufio"
	"fmt"
	"os"
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
	lines := os.Args[1:]
	
	for _, line := range lines {
		err := handleLine(line)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func handleLine(task string) error {
	fmt.Print("Полученная задача на обработку: ", task, "\n")
	return nil
}

func handleStdin() error {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		line := scanner.Text()
		err := handleLine(line)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
