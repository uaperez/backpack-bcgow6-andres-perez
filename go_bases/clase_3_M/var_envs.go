package main

import (
	"fmt"
	"os"
)

func main() {
	// La función en una ocasión de fallo puede retornarnos un error, es bueno analizarlo
	err := os.Setenv("NAME", "Juancho")
	if err != nil {
		panic(err.Error())
	}

	// Obtiene el valor de una variable de entorno a partir de su llave. Si no se encuentra, retorna un string vacío.
	name := os.Getenv("NAME")
	fmt.Println(name)

	/*
		El uso de LookupEnv es para determinar si la variable existe o no en la predefinición del sistema.
		En caso de existir, nos retorna el nombre de la variable.
		En caso de no existir, nos devolverá una bandera que nos avisará de ello.
	*/
	value, ok := os.LookupEnv("NAME")
	if !ok {
		panic("No se ha encontrado la variable de entorno")
	}

	fmt.Println("El resultado de la variable de entorno es:", value)
}
