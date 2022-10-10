package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductRequest struct {
	Id int
}

var products []ProductRequest

func main() {
	route := gin.Default()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	products := route.Group("/products")
	{
		products.POST("/")
	}
}

func Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token") // obtener un token de validacion
		if token != "123456" || token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token invalid",
			})
			return
		}

		var req ProductRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		req.Id = len(products) + 1

	}
}
