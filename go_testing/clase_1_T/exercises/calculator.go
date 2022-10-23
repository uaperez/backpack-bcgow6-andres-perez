package exercises

import "errors"

func Add(number1, number2 int) int {
	return number1 + number2
}

func Sub(number1, number2 int) int {
	return number1 - number2
}

func Divide(number1, number2 int) (int, error) {
	if number2 == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return number1 / number2, nil
}
