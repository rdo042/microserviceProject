package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_client"
)

type WebClientHandler struct {
	CreateClientUseCase create_client.CreateClientUseCase
}

func NewWebClientHandler(createClientUsecase create_client.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		CreateClientUseCase: createClientUsecase,
	}
}

func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto create_client.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	fmt.Println("X")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("X1")

	output, err := h.CreateClientUseCase.Execute(dto)
	fmt.Println(err)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("X2")

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("X3")
	w.WriteHeader(http.StatusCreated)
}
