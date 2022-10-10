package main

import (
	"errors"
	"fmt"
)

const (
	MINIMUM = "minimum"
	AVERAGE = "average"
	MAXIMUM = "maximum"
)

func calculateMin(numbers ...int) int {
	minimum := numbers[len(numbers)-1]
	for _, v := range numbers {
		if v < minimum {
			minimum = v
		}
	}
	return minimum
}

func calculateMax(numbers ...int) int {
	maximum := numbers[len(numbers)-1]
	for _, v := range numbers {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}

func calculateAvg(numbers ...int) int {
	var average int
	for _, v := range numbers {
		average += v
	}
	return average / len(numbers)
}

func generateOperation(operation string) (func(numbers ...int) int, error) {
	switch operation {
	case MINIMUM:
		return calculateMin, nil
	case MAXIMUM:
		return calculateMax, nil
	case AVERAGE:
		return calculateAvg, nil
	}
	return nil, errors.New(fmt.Sprintf("La operación '%s' no está definida en el programa", operation))
}

func main() {
	minFunc, _ := generateOperation(MINIMUM)
	maxFunc, _ := generateOperation(MAXIMUM)
	averageFunc, _ := generateOperation(AVERAGE)

	minValue := minFunc(2, 4, 5, 6, 1, 24)
	maxValue := maxFunc(2, 4, 5, 6, 1, 24)
	avgValue := averageFunc(2, 4, 5, 6, 1, 24)

	fmt.Printf("Min: %v | Max: %v | Average: %v\n", minValue, maxValue, avgValue)
}
