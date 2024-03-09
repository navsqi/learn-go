package main

import (
	"errors"
	"fmt"
)

// function

func main() {
	printMe("Hello world")
	result, remainder, err := intDivision(11, 0)

	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}

	fmt.Println()

	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

// (numerator int, denominator int) ==> define parameter untuk function
// (int, int, error) ==> define tipedata response, jika bukan function void
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cant divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
