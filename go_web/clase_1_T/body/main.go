package main

type Employee struct {
	Name     string
	Password string
	Id       int
	IsActive bool `form:active`
}

func main() {

}
