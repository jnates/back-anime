package types

import "backend_crudgo/domain/products/domain/model"

//ProductCreateResponse to message for response handler products.
type ProductCreateResponse struct {
	Message string `json:"message,omitempty"`
}

type ProductResponse struct {
	Message string        `json:"message"`
	Product *model.Product `json:"product,omitempty"`
	Error   string        `json:"error,omitempty"`
}

//ProductGetResponse Struct response product.
type ProductGetResponse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Amount       int    `json:"amount"`
	UserCreated  string `json:"user_created"`
	DateCreated  string `json:"date_created"`
	UserModified string `json:"user_modified"`
	DateModified string `json:"date_modified"`
}

