package main

import "fmt"

func main() {
	age := 23
	isEmployed := true
	yearsOfWork := 1.5
	salary := 100_001

	if age <= 22 && !isEmployed && yearsOfWork <= 1 {
		fmt.Println("No cumples con las condiciones para que te podamos prestar...")
	} else {
		fmt.Println("Wepa! Cumples con los requisitos para poder obtener el préstamo!")
		if salary > 100_000 {
			fmt.Println("Inclusive podrás tener el beneficio de no tener interes, genial!!")
		} else {
			fmt.Println("[!] Te cobraremos una tasa de interés, okey?")
		}
	}
}
