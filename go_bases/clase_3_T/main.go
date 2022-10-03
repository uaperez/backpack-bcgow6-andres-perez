package main

import "fmt"

func main() {
	var pointer *int
	pointer2 := new(int)
	fmt.Println(pointer, pointer2)

	var number int = 10
	pointer3 := &number

	fmt.Printf("El valor del número es %d\n", number)

	Incrementar(pointer3)

	fmt.Printf("El valor del número es %d\n", number)

	var c = Coso{Valor: 18}

	fmt.Printf("%+v", c)
	c.Actualizar(15)

	fmt.Printf("%+v", c)

}

func Incrementar(p *int) {
	*p = 20
}

type Coso struct {
	Valor int
}

func (c *Coso) Actualizar(new int) {
	c.Valor = new
}
