package persistence

import (
	"backend_crudgo/domain/products/domain/model"
	repoDomain "backend_crudgo/domain/products/domain/repository"
	"backend_crudgo/infrastructure/database"
	response "backend_crudgo/types"
	"context"
	"database/sql"
)

type sqlProductRepo struct {
	Conn *database.DataDB
}

//NewProductRepository Should initialize the dependencies for this service.
func NewProductRepository(Conn *database.DataDB) repoDomain.ProductRepository {
	return &sqlProductRepo{
		Conn: Conn,
	}
}

func (sr *sqlProductRepo) CreateProductHandler(ctx context.Context, product *model.Product) (*response.ProductCreateResponse, error) {
	stmt, err := sr.Conn.DB.PrepareContext(ctx, InsertProduct)
	if err != nil {
		return &response.ProductCreateResponse{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, &product.ProductID, &product.ProductName, &product.ProductAmount, &product.ProductUserCreated, &product.ProductUserModify)
	var idResult string
	err = row.Scan(&idResult)
	if err != sql.ErrNoRows {
		return &response.ProductCreateResponse{}, err
	}
	ProductResponse := response.ProductCreateResponse{
		Message: "Product created",
	}
	return &ProductResponse, nil
}

func (sr *sqlProductRepo) GetProductHandler(ctx context.Context, id string) (*response.ProductResponse, error) {
	stmt, err := sr.Conn.DB.PrepareContext(ctx, SelectProduct)
	if err != nil {
		return &response.ProductResponse{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)
	product := &model.Product{}
	err = row.Scan(&product.ProductID, &product.ProductName, &product.ProductAmount, &product.ProductUserCreated, &product.ProductDateCreated, &product.ProductUserModify,
		&product.ProductDateModify)
	if err != nil {
		return &response.ProductResponse{Error: err.Error()}, err
	}

	productResponse := &response.ProductResponse{
		Message: "Get product success",
		Product: product,
	}

	return productResponse, nil
}
