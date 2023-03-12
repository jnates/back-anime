package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"backend_crudgo/domain/products/domain/model"
	repoDomain "backend_crudgo/domain/products/domain/repository"
	"backend_crudgo/domain/products/infrastructure/persistence"
	"backend_crudgo/infrastructure/database"
	"backend_crudgo/infrastructure/middleware"
)

//ProductRouter router
type ProductRouter struct {
	Repo repoDomain.ProductRepository
}

//NewProductHandler Should initialize the dependencies for this service.
func NewProductHandler(db *database.DataDB) *ProductRouter {
	return &ProductRouter{
		Repo: persistence.NewProductRepository(db),
	}
}

// CreateProductHandler Created initialize handler product.
func (prod *ProductRouter) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	ctx := r.Context()
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusBadRequest, "bad request", err.Error())
		return
	}
	result, err := prod.Repo.CreateProductHandler(ctx, &product)
	if err != nil {
		_ = middleware.HTTPError(w, r, http.StatusConflict, "Conflict", err.Error())
		return
	}
	w.Header().Add("Location", fmt.Sprintf("%s%s", r.URL.String(), result))
	_ = middleware.JSON(w, r, http.StatusCreated, result)
}

func (prod *ProductRouter) GetProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		productResponse, err := prod.Repo.GetProductHandler(r.Context(), id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		jsonBytes, err := json.Marshal(productResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}
}
