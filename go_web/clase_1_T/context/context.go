package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/saludo", func(ctx *gin.Context) {
		request := ctx.Request
		fmt.Println(request)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hola",
		})
	})

	router.Run()
}
