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
		h.logger.Info("could not get user_id from context")
		resp = response.BadRequest
		return
	}
	var input models.Pagination
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		h.logger.Error("binding error: " + err.Error())
		resp = response.BadRequest
		resp.Message = err.Error()
		return
	}
	trns, err := h.svc.GetMonthlyTrns(userId, input)
	if err != nil {
		resp = response.InternalServer
		resp.Message = err.Error()
		return
	}
	resp = response.Success
	resp.Payload = trns
}
