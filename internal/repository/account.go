package repository

import (
	"wallets/internal/models"
	e "wallets/pkg/errors"
)

func (repo *repository) RefillBalance(amount int64, accId int64, user models.UserForBalance) (err error) {
	var (
		acc models.Account
	)
	tx := repo.DB.Begin()
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

func (repo *repository) GetAccById(id int64) (acc models.Account, err error) {
	err = repo.DB.Find(&acc, id).Error
	if err != nil {
		err = e.ErrAccNotFound
		return
	}
	return
}

func (repo *repository) GetBalance(accId int64) (balance int64, err error) {
	err = repo.DB.Raw("SELECT balance FROM accounts WHERE id = ?", accId).Scan(&balance).Error
	if err != nil {
		return
	}
	return
}
