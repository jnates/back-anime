package repository

import (
	"backend_crudgo/domain/products/domain/model"
	response "backend_crudgo/types"
	"context"
)

//ProductRepository interfaces handlers products
type ProductRepository interface {
	CreateProductHandler(ctx context.Context, product *model.Product) (*response.ProductCreateResponse, error)
}

//GetAllProducts type
type GetAllProducts interface {
	GetALLHandler(ctx context.Context, product *model.Product) (*response.ProductALLResponse, error)
}
