package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Meenu")

	var numerator int = 11
	var denominator int = 2
	var result, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("The quotient is %v\n", result)
}

func printMe(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func intDivision(numerator int, denominator int) (int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, err
	}
	var result int = numerator / denominator
	return result, nil
}
