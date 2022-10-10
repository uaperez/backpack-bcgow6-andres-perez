package main

import (
	"errors"
	"fmt"
)

const (
	DOG = "dog"
	CAT = "cat"
)

func dogFunc(count int) int {
	return count * 10
}

func catFunc(count int) int {
	return count * 5
}

func Animal(typeAnimal string) (func(count int) int, error) {
	switch typeAnimal {
	case DOG:
		return dogFunc, nil
	case CAT:
		return catFunc, nil
	}
	return nil, errors.New("Este animal no está disponible!")
}

func main() {
	dogFunc, _ := Animal(DOG)
	catFunc, _ := Animal(CAT)

	fmt.Println(dogFunc(15))
	fmt.Println(catFunc(10))
}
