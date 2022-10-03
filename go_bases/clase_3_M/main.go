package main

import "fmt"

func main() {
	name := "Juancho"
	age := 19

	/*
		%10s => significa que antecederá diez espacios antes del inicio del valor de la variable que se va a formatear
		Si lo almacenásemos en una variable, nos devolverá una cantidad de bytes y un error.
	*/
	fmt.Printf("Hola, %10s! Tienes %d años.\n", name, age)

	/*
		Sprint permite retornar un texto formateado en vez de imprimirlo.
		Tiene sus variantes de la familia Println, Print y Printf.

		Sprintf recibe como primera parámetro un string que hace par a la plantilla y como segundo parámetro recibe un ellipsis de datos a formatear en la plantilla.
	*/
	text := fmt.Sprintf("Hola, %10s! Tienes %d años.\n", name, age)

	fmt.Println("Tu variable de texto formateado es:", text)
}
