package repository

import "wallets/internal/models"

func (repo *Repository) GetUserByPhone(phone string) (user models.User, err error) {
	err = repo.DB.Raw("select * from users").Scan(&user).Error
	if err != nil {
		return user, err
	}
	return
}
