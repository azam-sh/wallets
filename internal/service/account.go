package service

import (
	"wallets/internal/models"
	e "wallets/pkg/errors"
)

func (s *service) TopUpBalance(input models.TopUpBalanceReq) (err error) {
	if input.AccountId < 1 || input.Amount <= 0 {
		s.logger.Info("account_id not specified or amount < 0")
		return e.ErrInvalidInput
	}
	user, err := s.repo.GetUserByAccId(input.AccountId)
	if err != nil {
		return
	}
	err = s.repo.TopUpBalance(input.Amount, input.AccountId, user)
	if err != nil {
		return
	}
	return
}

func (s *service) GetBalance(accId int64) (balance int64, err error) {
	if accId < 1 {
		s.logger.Info("account_id not specified")
		err = e.ErrInvalidInput
		return
	}
	balance, err = s.repo.GetBalance(accId)
	if err != nil {
		return
	}
	return
}
