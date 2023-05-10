package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Breil11/calculator/pkg/calculator"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println("Usage: calculator <number1> <operation> <number2>")
		os.Exit(1)
	}

	number1, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid number1")
		os.Exit(1)
	}

	number2, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid number2")
		os.Exit(1)
	}

	var result int

	switch args[1] {
	case "+":
		result = calculator.Add(number1, number2)
	case "-":
		result = calculator.Subtract(number1, number2)
	case "*":
		result = calculator.Multiply(number1, number2)
	case "/":
		result = calculator.Divide(number1, number2)
	default:
		fmt.Println("Invalid operation")
		os.Exit(1)
	}

	fmt.Println(result)
}
