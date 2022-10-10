package main

import "fmt"

type product struct {
	Type  string
	Name  string
	Price float64
}

type Product interface {
	CalculatePrice() float64
}

func (p *product) CalculatePrice() float64 {
	switch p.Type {
	case "Small":
		return p.Price
	case "Medium":
		return p.Price + (0.3 * p.Price)
	case "Big":
		return p.Price + (0.3 * p.Price) + 2500
	}
	return 0
}

type Ecommerce interface {
	Total() float64
	Add(product Product)
}

type store struct {
	Products []Product
}

func (s *store) Total() float64 {
	var totalPrice float64
	for _, v := range s.Products {
		totalPrice += v.CalculatePrice()
	}
	return totalPrice
}

func (s *store) Add(product Product) {
	s.Products = append(s.Products, product)
}

func NewProduct(typeProduct, name string, price float64) Product {
	return &product{typeProduct, name, price}
}

func NewStore() Ecommerce {
	return &store{}
}

func main() {

	myStore := NewStore()
	product1 := NewProduct("Big", "Machine Box", 1456.0)
	product2 := NewProduct("Small", "Mousepad", 120.2)
	myStore.Add(product1)
	myStore.Add(product2)

	fmt.Println(myStore.Total())

}
