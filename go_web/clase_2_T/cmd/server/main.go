package server

import (
	"github.com/gin-gonic/gin"
	"github.com/uaperez/backpack-bcgow6-andres-perez/go_web/clase_2_T/cmd/server/handler"
	"github.com/uaperez/backpack-bcgow6-andres-perez/go_web/clase_2_T/internal/products"
)

func main() {
	repository := products.NewRepository()
	service := products.NewService(repository)
	productRoute := handler.NewProduct(service)

	router := gin.Default()

	pr := router.Group("/products")
	{
		pr.GET("/", productRoute)
	}
}
