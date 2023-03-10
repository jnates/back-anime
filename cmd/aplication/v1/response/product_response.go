package response

//ProductCreateResponse to message for response handler products.
type ProductCreateResponse struct {
	Message string `json:"message,omitempty"`
}

//ProductALLResponse Struct response product.
type ProductALLResponse struct {
	ProductName       string `json:"product_name,omitempty"`
	Amount             string `json:"amount,omitempty"`
	ProductUserCreated string `json:"product_user_created,omitempty"`
}
