package repository

import (
	response "backend_crudgo/cmd/aplication/v1/response"
	"backend_crudgo/domain/products/domain/model"
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
