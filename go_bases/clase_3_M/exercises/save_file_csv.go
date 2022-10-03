package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	Id    int64
	Price float64
	Stock int64
}

func (p Product) ToCsv() string {
	productAsString := fmt.Sprintf("%v;%f;%v", p.Id, p.Price, p.Stock)
	return productAsString
}

func generateProducts() []Product {
	var products []Product
	products = append(products, Product{1, 1970.23, 4})
	products = append(products, Product{2, 894.4, 5})
	products = append(products, Product{3, 7874.45, 2})
	return products
}

func generateCsvFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err.Error())
	}

	file.WriteString("Id;Price;Stock\n")

	for index, product := range generateProducts() {
		line := product.ToCsv()
		if index != (len(generateProducts()) - 1) {
			line += "\n"
		}
		file.WriteString(line)
	}
}

func printCsvFile(filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}

	lines := strings.Split(string(file), "\n")

	headers := strings.Split(lines[0], ";")
	products := lines[1:]

	fmt.Printf("%v %20v %4v\n", headers[0], headers[1], headers[2])
	var totalPriceProducts float64

	for _, v := range products {
		infoProduct := strings.Split(strings.Trim(v, "\n"), ";")
		fmt.Printf("%v %20v %4v\n", infoProduct[0], infoProduct[1], infoProduct[2])
		priceAsFloat, err := strconv.ParseFloat(infoProduct[1], 8)
		if err != nil {
			panic(err.Error())
		}
		totalPriceProducts += priceAsFloat
	}

	fmt.Printf("\tTotal: %f\n", totalPriceProducts)

}

func main() {
	filename := "products.csv"
	generateCsvFile(filename)

	printCsvFile(filename)
}
