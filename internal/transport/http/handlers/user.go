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
		h.logger.Error("binding error: " + err.Error())
		resp = response.BadRequest
		resp.Message = err.Error()
		return
	}

	accounts, err := h.svc.CheckAccount(input.Phone)
	if err != nil {
		resp = response.InternalServer
		resp.Message = err.Error()
		return
	}

	resp = response.Success
	resp.Payload = accounts
}
