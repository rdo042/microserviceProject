package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com.br/rdo042/goappbalance/internal/usecase/find_balance"
)

type WebBalanceHandler struct {
	FindBalanceUseCase find_balance.FindBalanceUseCase
}

func NewWebBalanceHandler(findBalanceUsecase find_balance.FindBalanceUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		FindBalanceUseCase: findBalanceUsecase,
	}
}

func (h *WebBalanceHandler) FindBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("E")

	var dto find_balance.FindBalanceInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	fmt.Println(err)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := h.FindBalanceUseCase.Execute(dto)
	fmt.Println(err)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
