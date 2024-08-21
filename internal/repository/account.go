package repository

import (
	"wallets/internal/models"
	e "wallets/pkg/errors"
)

func (r *repository) TopUpBalance(amount int64, accId int64, user models.UserForBalance) (err error) {
	var (
		acc models.Account
	)
	tx := r.db.Begin()
	err = tx.Raw("SELECT FOR UPDATE * FROM accounts WHERE id = ?", accId).Scan(&acc).Error
	if err != nil {
		tx.Rollback()
		return
	}
	if acc.Id < 1 {
		tx.Rollback()
		return e.ErrAccNotFound
	}

	newBalance := amount + acc.Balance

	if newBalance > user.MaxBalance {
		tx.Rollback()
		return e.ErrExceededLimit
	}

	trn := models.Transaction{
		Amount:        amount,
		AccountId:     acc.Id,
		UserId:        user.Id,
		BalanceBefore: acc.Balance,
		BalanceAfter:  newBalance,
	}

	err = tx.Create(&trn).Error
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Exec("UPDATE accounts SET balance = ? WHERE id = ?", newBalance, acc.Id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	return nil
}

func (r *repository) GetAccById(id int64) (acc models.Account, err error) {
	err = r.db.Find(&acc, id).Error
	if err != nil {
		err = e.ErrAccNotFound
		return
	}
	return
}

func (r *repository) GetBalance(accId int64) (balance int64, err error) {
	err = r.db.Raw("SELECT balance FROM accounts WHERE id = ?", accId).Scan(&balance).Error
	if err != nil {
		return
	}
	return
}
