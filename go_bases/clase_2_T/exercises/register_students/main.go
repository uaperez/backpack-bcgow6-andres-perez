package main

import "fmt"

type Student struct {
	IdCard     string
	FirstName  string
	LastName   string
	DateJoined string
}

func (s *Student) GetDetails() {
	fmt.Println("Nombre:", s.FirstName)
	fmt.Println("Apellido:", s.LastName)
	fmt.Println("DNI:", s.IdCard)
	fmt.Println("Fecha ingreso:", s.DateJoined)
	fmt.Println("===============================")
}

func main() {
	students := []Student{
		{"1234567890", "Juan Andrés", "Pérez", "19/09/2022"},
		{"0987654321", "Laura", "Pérez", "12/02/2004"},
	}

	fmt.Println("============== STUDENTS INFO =================")
	for _, v := range students {
		v.GetDetails()
	}

}
