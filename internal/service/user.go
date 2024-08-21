package service

import (
	"wallets/internal/models"
	e "wallets/pkg/errors"
)

func (s *service) CheckAccount(phone string) (acc models.CheckAccResp, err error) {
	if phone == "" || len(phone) != 9 {
		err = e.ErrIncorrectPhoneNumber
		return
	}

	acc, err = s.repo.GetAccByPhone(phone)
	if err != nil {
		return
	}

	return
}
