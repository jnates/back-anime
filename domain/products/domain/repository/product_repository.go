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

//// GetAllProducts es la interfaz que define el método para obtener todos los productos.
//type GetAllProducts interface {
//	GetAllProductsHandler(ctx context.Context) (*response.ProductGetResponse, error)
//}