package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env", err)
	}
	usuario := os.Getenv("MY_USER")
	password := os.Getenv("MY_PASS")
	fmt.Println(usuario, password)
}
