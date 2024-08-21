package service

import "wallets/internal/models"

func (s *service) GetMonthlyTrns(userId int64, input models.Pagination) (trns models.TrnsHistory, err error) {
	if input.Page <= 0 {
		input.Page = 1
	}
	switch {
	case input.Rows > 30:
		input.Rows = 30
	case input.Rows <= 0:
		input.Rows = 15
	}

	trns, err = s.Repo.GetMonthlyTrns(userId, input)
	if err != nil {
		return
	}
	return
}
