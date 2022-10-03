package main

import (
	"fmt"
	"os"
)

func main() {
	message := "Hello world, I'm a new booooooootcamper :)!"
	err := os.WriteFile("./test.txt", []byte(message), 0644)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Texto escrito correctamente!")
	os.Exit(1)

	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Entradas recolectadas hasta que ocurrió el error:", files)
		panic(err.Error())
	}

	for _, value := range files {
		fmt.Println(value.Name())
		info, _ := value.Info()
		info.IsDir()
		info.Name()
	}

	fileContentAsBytes, err := os.ReadFile("./main.go")
	if err != nil {
		panic(err.Error())
	}

	/*
		%s => es para que pase de bytes a string
		En caso de no colocarse su formatter, simplemente nos dará un array de bytes
	*/
	fmt.Printf("%s", fileContentAsBytes)
	fmt.Println(string(fileContentAsBytes))
	fmt.Println(fileContentAsBytes)
}
