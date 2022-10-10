package main

import (
	"errors"
	"fmt"
)

func calculateAverage(numbers ...int) (int, error) {
	var average int
	for _, v := range numbers {
		if v < 0 {
			return 0, errors.New(fmt.Sprintf("Se ha encontrado un número negativo: %v", v))
		}
		average += v
	}
	return (average / len(numbers)), nil
}

func main() {
	average, err := calculateAverage(1, 3, 4, 5, 2, 5, 3, 1)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("El promedio de las calificaciones es de", average)
}
