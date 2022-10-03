package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var (
	FILENAME_CUSTOMERS = "customers.txt"
)

type Customer struct {
	FileId      int
	FirstName   string
	LastName    string
	IdCard      string
	PhoneNumber string
	Address     string
}

type Company struct {
	Name      string
	Customers []Customer
}

func getFile(filename string) string {
	defer func() {
		/*err := recover()
		if err != nil {
			fmt.Println("[ERROR]", err)
		}*/
		fmt.Println("Ejecución finalizada")
	}()
	file, err := os.ReadFile(filename)
	if err != nil {
		panic("El archivo indicado no fue encontrado o está dañado")
	}
	return string(file)
}

func generateCustomerFileId() {
	numberId := rand.Int()
	arr := []byte(strconv.Itoa(numberId))
	fmt.Println(numberId)
	fmt.Println(string(arr[1]))
}

/*func (c *Company) SaveUser(customer Customer) {
	get
}*/

func main() {

	//getFile("customers.txt")

	generateCustomerFileId()
	generateCustomerFileId()
	generateCustomerFileId()

	fmt.Println("Hola probando")

}
