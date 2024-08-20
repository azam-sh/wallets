package handlers

import (
	"net/http"
	"wallets/pkg/response"
)

func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	response := response.Response{
		Code:    http.StatusOK,
		Message: "pong",
	}
	response.WriteJSON(w)
}
