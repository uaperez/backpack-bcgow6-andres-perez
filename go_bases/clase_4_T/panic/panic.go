package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Hey!! Ocurrió un panic :O ", err)
		}
		fmt.Println("Saliendo bajo todo pronóstico :D")
	}()

	var a *int
	b := 10
	a = &b
	fmt.Println(a)

	_, err := os.ReadFile("this file don't exists")
	if err != nil {
		panic(err)
	}

	fmt.Println("Todo bien, todo correcto! :D")
}
