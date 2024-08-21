package repository

import (
	"wallets/internal/models"
	e "wallets/pkg/errors"
)

func (r *repository) GetAccByPhone(phone string) (acc models.CheckAccResp, err error) {
	err = r.db.Raw("SELECT id, first_name, last_name, identified FROM users WHERE phone = ?", phone).Scan(&acc).Error
	if err != nil {
		r.logger.Error("get user info by phone sql err: " + err.Error())
		return
	}
	if acc.Id < 1 {
		r.logger.Info("user by phone not found")
		err = e.ErrAccNotFound
		return
	}
	err = r.db.Raw("SELECT * FROM accounts WHERE user_id = ?", acc.Id).Scan(&acc.Accounts).Error
	if err != nil {
		r.logger.Error("get account by phone sql err: " + err.Error())
		return
	}
	if len(acc.Accounts) < 1 {
		r.logger.Info("account by phone not found")
		err = e.ErrAccNotFound
		return
	}
	return
}

func (r *repository) GetUserByAccId(id int64) (user models.UserForBalance, err error) {
	err = r.db.Raw(`
		SELECT u.id, identified, phone, max_balance FROM users u
		LEFT JOIN balance_limits bl ON u.limit_id = bl.id
		WHERE u.id = (SELECT a.user_id FROM accounts a WHERE a.id = ?)`, id).Scan(&user).Error
	if err != nil {
		r.logger.Error("get user by account_id sql err: " + err.Error())
		return
	}
	if user.Id < 1 {
		r.logger.Info("user by account_id not found")
		err = e.ErrAccNotFound
		return
	}
	return
}
