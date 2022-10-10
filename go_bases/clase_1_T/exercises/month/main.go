package main

import "fmt"

func main() {
	// Forma 1, con un mapa
	months := map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	selectedMonth := 12
	month := months[selectedMonth]

	if month != "" {
		fmt.Printf("%v, %s\n", selectedMonth, month)
	} else {
		fmt.Println("This month don't exists")
	}

	// Forma dos, con ifs anidados
	if selectedMonth == 1 {
		fmt.Printf("%v, Enero\n", selectedMonth)
	} else if selectedMonth == 2 {
		fmt.Printf("%v, Febrero\n", selectedMonth)
	} else if selectedMonth == 3 {
		fmt.Printf("%v, Marzo\n", selectedMonth)
	} else if selectedMonth == 4 {
		fmt.Printf("%v, Abril\n", selectedMonth)
	} else if selectedMonth == 5 {
		fmt.Printf("%v, Mayo\n", selectedMonth)
	} else if selectedMonth == 6 {
		fmt.Printf("%v, Junio\n", selectedMonth)
	} else if selectedMonth == 7 {
		fmt.Printf("%v, Julio\n", selectedMonth)
	} else if selectedMonth == 8 {
		fmt.Printf("%v, Agosto\n", selectedMonth)
	} else if selectedMonth == 9 {
		fmt.Printf("%v, Septiembre\n", selectedMonth)
	} else if selectedMonth == 10 {
		fmt.Printf("%v, Octubre\n", selectedMonth)
	} else if selectedMonth == 11 {
		fmt.Printf("%v, Noviembre\n", selectedMonth)
	} else if selectedMonth == 12 {
		fmt.Printf("%v, Diciembre\n", selectedMonth)
	} else {
		fmt.Println("This month don't exists")
	}
}
