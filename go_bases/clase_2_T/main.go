package main

import "fmt"

type Perro struct {
	Nombre string
	Raza   string
	Peso   float64
	Color  []string
}

func (p Perro) Ladrar() {
	fmt.Println("guauu!!")
}

type Gato struct {
	Nombre   string  `json:"nombre"` // etiquetas
	Edad     int     `json:"edad"`
	Peso     float64 `json:"peso"`
	Raza     string  `json:"raza"`
	Caracter string  `json:"caracter"`
}

func (g Gato) Dormir() {
	fmt.Println("Estoy durmiendo...")
}

func (g Gato) Comer() {
	fmt.Println("Hola, estoy comiendo, miaaau")
}

type Animal interface {
	Dormir()
	Comer()
}

func NewAnimal() Animal {
	return &Gato{"michi", 12, 12.9, "antonio", "dormilon"}
}

func main() {
	cat := NewAnimal()
	cat.Comer()
	fmt.Println("Hola")
}
