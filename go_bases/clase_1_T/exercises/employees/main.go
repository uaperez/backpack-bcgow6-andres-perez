package main

import "fmt"

func main() {
	var employees = map[string]int{
		"Benjamin": 20,
		"Nahuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	// Punto 1
	fmt.Println("La edad de benjamin es de", employees["Benjamin"])

	// Punto 2
	countEmployeesMajors := 0
	for _, ageOfEmployee := range employees {
		if ageOfEmployee > 21 {
			countEmployeesMajors++
		}
	}

	if countEmployeesMajors > 0 {
		fmt.Println("La cantidad de empleados mayores de 21 son", countEmployeesMajors)
	} else {
		fmt.Println("No tienes empleados mayores de 21 años")
	}

	// Punto 3
	employees["Federico"] = 25

	// Punto 4
	delete(employees, "Pedro")
}
