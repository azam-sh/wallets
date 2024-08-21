package handlers

import (
	"encoding/json"
	"net/http"
	"wallets/internal/models"
	"wallets/pkg/response"

	"github.com/gorilla/context"
)

func (h *Handler) GetMonthlyTrns(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	userId, ok := context.Get(r, "userID").(int64)
	if !ok {
		resp = response.BadRequest
		return
	}
	var input models.Pagination
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		resp = response.BadRequest
		return
	}
	trns, err := h.svc.GetMonthlyTrns(userId, input)
	if err != nil {
		resp = response.InternalServer
		return
	}

	resp.Payload = trns
	resp = response.Success
}
