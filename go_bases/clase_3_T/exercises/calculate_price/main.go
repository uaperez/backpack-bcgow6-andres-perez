package main

import (
	"fmt"
	"math"
)

type Company struct {
	Products     []ProductItem
	Services     []Service
	Maintenances []Maintenance
}

type ProductItem struct {
	Name  string
	Price float64
	Stock uint16
}

type Service struct {
	Name          string
	Price         float64
	WorkedMinutes uint16
}

type Maintenance struct {
	Name  string
	Price float64
}

type Calculable interface {
	CalculatePrice() float64
}

func (pi *ProductItem) CalculatePrice() float64 {
	return pi.Price * float64(pi.Stock)
}

func (s *Service) CalculatePrice() float64 {
	times := s.WorkedMinutes / 30
	return s.Price * float64(math.Round(float64(times)))
}

func (m *Maintenance) CalculatePrice() float64 {
	return m.Price
}

// TODO: Implementar patrón Composite? Maybe.

func (c *Company) AddProducts(products []ProductItem) {
	c.Products = products
}

func (c *Company) AddServices(services []Service) {
	c.Services = services
}

func (c *Company) AddMaintenances(maintenances []Maintenance) {
	c.Maintenances = maintenances
}

func (c *Company) CalculateTotalPriceOfProducts(channel chan float64) {
	var totalPrice float64
	for _, product := range c.Products {
		totalPrice += product.CalculatePrice()
	}
	fmt.Printf("Total of products: %.2f\n", totalPrice)
	channel <- totalPrice
}

func (c *Company) CalculateTotalPriceOfServices(channel chan float64) {
	var totalPrice float64
	for _, service := range c.Services {
		totalPrice += service.CalculatePrice()
	}
	fmt.Printf("Total of services: %.2f\n", totalPrice)
	channel <- totalPrice
}

func (c *Company) CalculateTotalPriceOfMaintenances(channel chan float64) {
	var totalPrice float64
	for _, maintenance := range c.Maintenances {
		totalPrice += maintenance.CalculatePrice()
	}
	fmt.Printf("Total of maintenances: %.2f\n", totalPrice)
	channel <- totalPrice
}

func (c *Company) CalculateCosts() float64 {
	channel := make(chan float64)
	go c.CalculateTotalPriceOfProducts(channel)

	go c.CalculateTotalPriceOfServices(channel)

	go c.CalculateTotalPriceOfMaintenances(channel)
	var lectura float64

	for i := 0; i < 3; i++ {
		lectura += <-channel
	}

	return lectura
}

func main() {

	products := []ProductItem{
		ProductItem{"PC Gamer", 3456.90, 4},
		ProductItem{"Impresora", 4000.50, 5},
	}

	services := []Service{
		Service{"Limpieza de PC", 40.50, 60},
		Service{"Instalacion de Office", 35.0, 45},
	}

	maintenances := []Maintenance{
		Maintenance{"Backup de Servidores", 4000.0},
		Maintenance{"Revisión de Firewall", 2000.0},
	}

	company := &Company{}

	company.AddProducts(products)
	company.AddServices(services)
	company.AddMaintenances(maintenances)

	costs := company.CalculateCosts()
	fmt.Printf("\nTotal ticket: %.1f\n", costs)

}
