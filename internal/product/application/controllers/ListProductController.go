package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ldez/go-monorepo-example/product/application/usecases"
)

type ListProductController struct {
	ListProductUseCase *usecases.ListProductUseCase
}

func (controller *ListProductController) Handle(w http.ResponseWriter, r *http.Request) error {
	products := controller.ListProductUseCase.Execute()
	response, jsonErr := json.Marshal(products)

	if jsonErr != nil {
		return jsonErr
	}

	w.Write(response)
	return nil
}
