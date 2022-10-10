package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Height    int    `json:"height"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt"`
}

func generateUsers() []User {
	users := []User{
		{1, "Juan", "Perez", "juanandres.perez@mercadolibre.com.co", 19, 176, true, "19-09-2022"},
		{2, "John", "Doe", "johndoe@enterprises.co", 50, 160, false, "14-04-2021"},
		{3, "Jane", "Doe", "jane.doe@noidentity.wal.es", 45, 167, true, "01-12-2019"},
	}
	return users
}

func GetAll(ctx *gin.Context) {
	// TODO: Falta el que es con filtrados
	users := generateUsers()
	ctx.JSON(http.StatusOK, gin.H{
		"total":   len(users),
		"results": users,
	})
}

func GetOne(ctx *gin.Context) {
	users := generateUsers()
	idUserSearched := ctx.Param("id")
	var userFound User
	var response gin.H
	for _, user := range users {
		id := strconv.Itoa(user.Id)
		if id == idUserSearched {
			userFound = user
			break
		}
	}
	if (User{} == userFound) {
		response = gin.H{
			"message": "User not found",
		}
		ctx.JSON(http.StatusNotFound, response)
	} else {
		response = gin.H{
			"user": userFound,
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func main() {
	router := gin.Default()
	router.GET("/users/:id", GetOne)
	router.GET("/users", GetAll)

	router.Run()
}
