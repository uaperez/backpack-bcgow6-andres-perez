package main

import (
	"errors"
	"fmt"
	"os"
)

func division(num1, num2 float64) (result float64, err error) {
	if num2 == 0 {
		err = errors.New("Cannot divide number by zero")
		return
	}

	result = num1 / num2
	return
}

func calculate(operator string, num1, num2 float64) (result float64, err error) {
	switch operator {
	case "/":
		result, err = division(num1, num2)
	default:
		err = errors.New("Operator unsupported")
	}
	return
}

func main() {
	result, err := calculate("/", 10, 2)
	if err != nil {
		fmt.Printf("Ha ocurrido un error inesperado: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
