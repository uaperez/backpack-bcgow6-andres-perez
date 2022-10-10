package main

import "fmt"

func calculateTax(salary float32) float32 {
	var taxes float32
	if salary > 50_000 {
		taxes = 0.17 * salary
	}
	if salary > 150_000 {
		taxes += 0.10 * salary
	}
	return taxes
}

func main() {
	fmt.Println("Por tu salario, deberás de pagar los siguientes impuestos:", calculateTax(50000.0))
	fmt.Println("Por tu salario, deberás de pagar los siguientes impuestos:", calculateTax(500000.0))
	fmt.Println("Por tu salario, deberás de pagar los siguientes impuestos:", calculateTax(54000.0))
}
