package handlers

import (
	"encoding/json"
	"net/http"
	"wallets/internal/models"
	"wallets/pkg/response"
)

func (h *Handler) TopUpBalance(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	var input models.TopUpBalanceReq

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		h.logger.Error("binding error: " + err.Error())
		resp = response.BadRequest
		resp.Message = err.Error()
		return
	}

	err = h.svc.TopUpBalance(input)
	if err != nil {
		resp = response.InternalServer
		resp.Message = err.Error()
		return
	}

	resp = response.Success
	resp.Message = "Баланс пополнен успешно!"
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	var resp response.Response
	defer resp.WriteJSON(w)

	var input models.BalanceReq

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		h.logger.Error("binding error: " + err.Error())
		resp = response.BadRequest
		resp.Message = err.Error()
		return
	}

	balance, err := h.svc.GetBalance(input.AccountId)
	if err != nil {
		resp = response.InternalServer
		resp.Message = err.Error()
		return
	}

	resp = response.Success
	resp.Payload = map[string]int64{"balance": balance}
}
