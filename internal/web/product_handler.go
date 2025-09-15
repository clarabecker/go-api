package web

import (
	"encoding/json"
	"net/http"

	"github.com/clarabecker/estudos-go/internal/usecases"
)


type ProjectHandlers struct{
	CreateProductUseCase *usecases.CreateProductUseCase
	ListProductsUseCase *usecases.ListProductsUseCase
}

func NewProductHandlers(
	createUC *usecases.CreateProductUseCase,
	listUC *usecases.ListProductsUseCase,
) *ProjectHandlers {
	return &ProjectHandlers{
		CreateProductUseCase: createUC,
		ListProductsUseCase:  listUC,
	}
}

func (p *ProjectHandlers) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var input usecases.CreateProductInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.CreateProductUseCase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ProjectHandlers) ListProductHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListProductsUseCase.Execute()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

