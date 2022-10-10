package main

import (
	"fmt"
	"os"
)

func main() {
	flag := false

	if flag {
		fmt.Println("Flag es true")
	} else {
		fmt.Println("Flag es false")
		os.Exit(1)
	}

	fmt.Println("Fin del código!!")
}
