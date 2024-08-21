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
	err = tx.Raw("SELECT * FROM accounts WHERE id = ? FOR UPDATE", accId).Scan(&acc).Error
	if err != nil {
		r.logger.Error("select for update sql err: " + err.Error())
		tx.Rollback()
		return
	}
	if acc.Id < 1 {
		r.logger.Info("account not found")
		tx.Rollback()
		return e.ErrAccNotFound
	}

	newBalance := amount + acc.Balance

	if newBalance > user.MaxBalance {
		r.logger.Info("balance limit exceeded")
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
		r.logger.Error("create trn sql err: " + err.Error())
		tx.Rollback()
		return
	}

	err = tx.Exec("UPDATE accounts SET balance = ? WHERE id = ?", newBalance, acc.Id).Error
	if err != nil {
		r.logger.Error("update account balance sql err: " + err.Error())
		tx.Rollback()
		return
	}

	tx.Commit()
	return nil
}

func (r *repository) GetBalance(accId int64) (balance int64, err error) {
	var exists bool
	err = r.db.Raw("SELECT EXISTS(SELECT 1 FROM accounts WHERE id = ?)", accId).Scan(&exists).Error
	if err != nil {
		r.logger.Error("select exists account sql err: " + err.Error())
		return
	}
	if !exists {
		r.logger.Info("account does not exist")
		err = e.ErrAccNotFound
		return
	}
	err = r.db.Raw("SELECT balance FROM accounts WHERE id = ?", accId).Scan(&balance).Error
	if err != nil {
		r.logger.Error("select account balance sql err: " + err.Error())
		return
	}
	return
}
