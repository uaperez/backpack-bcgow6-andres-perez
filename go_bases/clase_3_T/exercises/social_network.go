package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FirstName string `json:"nombres"`
	LastName  string `json:"apellidos"`
	Age       uint8  `json:"edad"`
	Email     string `json:"correo"`
	Password  string `json:"contraseña"`
}

func (u *User) ChangeName(firstName string, lastName string) {
	u.FirstName = firstName
	u.LastName = lastName
}

func (u *User) ChangeAge(age uint8) {
	u.Age = age
}

func (u *User) ChangeEmail(email string) {
	u.Email = email
}

func (u *User) ChangePassword(password string) {
	u.Password = password
}

func (u *User) ToJson() string {
	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		panic(err.Error())
	}
	return string(out)
}

func main() {

	user := &User{"Juan", "Andrés", 19, "juanandres.perez@mercadolibre.com.co", "1234"}

	user.ChangeAge(20)

	fmt.Println(user.ToJson())

}
