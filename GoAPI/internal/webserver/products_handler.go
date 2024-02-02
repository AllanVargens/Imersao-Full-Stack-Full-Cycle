package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/imersao17/goapi/internal/entity"
	"github.com/devfullcycle/imersao17/goapi/internal/service"
	"github.com/go-chi/chi/v5"
)

type WebProductHandler struct {
	ProductServer *service.ProductService
}

func NewWebProductHandler(productServer *service.ProductService) *WebProductHandler {
	return &WebProductHandler{ProductServer: productServer}
}

// Metodos

func (wph *WebProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	product, err := wph.ProductServer.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (wph *WebProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	product, err := wph.ProductServer.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (wph *WebProductHandler) GetProductByCategoryID(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")
	products, err := wph.ProductServer.GetProductsByCategoryID(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (wph *WebProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product *entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	response, err := wph.ProductServer.CreateProduct(product.Name, product.Description, product.CategoryID, product.ImageURL, product.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}
