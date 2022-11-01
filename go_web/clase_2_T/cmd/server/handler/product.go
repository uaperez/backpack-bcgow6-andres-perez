package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/uaperez/backpack-bcgow6-andres-perez/go_web/clase_2_T/internal/products"
)

type request struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Count int    `json:"count"`
	Price int    `json:"price"`
}

type Product struct {
	service products.Service
}

func NewProduct(service products.Service) *Product {
	return &Product{service: service}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
