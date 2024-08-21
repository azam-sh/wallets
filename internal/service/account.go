package service

import (
	"wallets/internal/models"
	e "wallets/pkg/errors"
)

func (s *service) RefillBalance(input models.RefillBalanceReq) (err error) {
	if input.AccountId < 1 || input.Amount <= 0 {
		return e.ErrInvalidInput
	}
	user, err := s.Repo.GetUserByAccId(input.AccountId)
	if err != nil {
		return
	}
	err = s.Repo.RefillBalance(input.Amount, input.AccountId, user)
	if err != nil {
		return
	}
	return
}

func (s *service) GetBalance(accId int64) (balance int64, err error) {
	if accId < 1 {
		err = e.ErrInvalidInput
		return
	}
	balance, err = s.Repo.GetBalance(accId)
	if err != nil {
		return
	}
	return
}
