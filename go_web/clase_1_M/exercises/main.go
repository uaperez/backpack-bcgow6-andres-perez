package main

import (
	"net/http"

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
	users := generateUsers()
	ctx.JSON(http.StatusOK, gin.H{
		"total":   len(users),
		"results": users,
	})
}

func main() {
	router := gin.Default()
	router.GET("/welcome", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "Bootcamper")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hola " + name,
		})
	})
	router.GET("/users", GetAll)

	router.Run()
}
