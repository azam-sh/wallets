package repository

import (
	"wallets/internal/models"
	e "wallets/pkg/errors"
)

func (r *repository) GetAccByPhone(phone string) (acc models.CheckAccResp, err error) {
	err = r.db.Raw("SELECT id, first_name, last_name, identified FROM users WHERE phone = ?", phone).Scan(&acc).Error
	if err != nil {
		return
	}
	if acc.Id < 1 {
		err = e.ErrAccNotFound
		return
	}
	err = r.db.Raw("SELECT * FROM accounts WHERE user_id = ?", acc.Id).Scan(&acc.Accounts).Error
	if err != nil {
		return
	}
	if len(acc.Accounts) < 1 {
		err = e.ErrAccNotFound
		return
	}
	return
}

func (r *repository) GetUserByAccId(id int64) (user models.UserForBalance, err error) {
	err = r.db.Raw(`
		SELECT id, identified, phone, max_balance FROM users u
		LEFT JOIN balance_limits bl ON u.limit_id = bl.id
		WHERE u.id = (SELECT user_id FROM accounts WHERE id = ?)`, id).Scan(&user).Error
	if err != nil {
		return
	}
	if user.Id < 1 {
		err = e.ErrAccNotFound
		return
	}
	return
}
