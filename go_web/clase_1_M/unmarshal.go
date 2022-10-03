package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name      string
	Price     float64
	Published bool
}

func main() {
	jsonData := []byte(`{"Name": "Macbook Pro", "Price": 900, "Published": true}`)
	var p Product

	err := json.Unmarshal(jsonData, &p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Producto:", p)
}
