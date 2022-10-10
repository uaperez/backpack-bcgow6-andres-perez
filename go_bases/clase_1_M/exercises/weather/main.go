package main

import "fmt"

func main() {
	var temperature int
	var humidity float32
	var pressure float32

	temperature = 15
	humidity = 14.5
	pressure = 2.5
	fmt.Printf("La humedad es de %v, junto a una temperatura de %v y una despresurización de %v", humidity, temperature, pressure)
}
