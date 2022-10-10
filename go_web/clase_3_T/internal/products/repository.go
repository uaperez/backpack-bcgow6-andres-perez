package products

import (
	"github.com/uaperez/backpack-bcgow6-andres-perez/go_web/clase_3_T/internal/domain"
	"github.com/uaperez/backpack-bcgow6-andres-perez/go_web/clase_3_T/pkg/store"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
}

type repository struct {
	db store.Store
}
