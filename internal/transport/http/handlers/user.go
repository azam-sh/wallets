package handlers

import (
	"encoding/json"
	"net/http"
	"wallets/internal/models"
	"wallets/pkg/response"
)

func (h *Handler) CheckAccount(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	var input models.CheckAccReq
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		resp = response.BadRequest
		return
	}

	accounts, err := h.svc.CheckAccount(input.Phone)
	if err != nil {
		resp = response.InternalServer
		resp.Message = err.Error()
		return
	}

	resp.Payload = accounts
	resp = response.Success
}
