package main

import "fmt"

type Client struct {
	FirstName string
	LastName  string
	Email     string
	Products  []Product
}

type Product struct {
	Name  string
	Price float32
	Stock int64
}

func CreateProduct(name string, price float32) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}

func AddProduct(product *Product, client *Client) {
	client.Products = append(client.Products, *product)
}

func DeleteAllProducts(client *Client) {
	client.Products = make([]Product, 0)
}

func main() {
	client := Client{
		FirstName: "Juan",
		LastName:  "Pérez",
		Email:     "juanandres.perez@mercadolibre.com.co",
	}

	fmt.Println(client)

	product1 := CreateProduct("PS5", 1260.40)

	AddProduct(product1, &client)

	fmt.Println(client)

	AddProduct(CreateProduct("Quitapelusas", 67.3), &client)

	fmt.Println(client)

	DeleteAllProducts(&client)

	fmt.Println(client)
}
