package repository

import (
	"math"
	"time"
	"wallets/internal/models"
)

func (r *repository) GetMonthlyTrns(userId int64, input models.Pagination) (resp models.TrnsHistory, err error) {
	var (
		count    int64
		dateFrom time.Time
		dateTo   time.Time
	)
	now := time.Now()
	dateTo = now
	dateFrom = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	query := r.db.Table("transactions").Where("created_at >= ? and created_at <= ? and user_id = ?", dateFrom, dateTo, userId)

	err = query.Count(&count).Error
	if err != nil {
		r.logger.Error("count of trns sql err: " + err.Error())
		return
	}

	offset := (input.Page - 1) * input.Rows

	err = query.Limit(input.Rows).Offset(offset).Find(&resp.Trns).Error
	if err != nil {
		r.logger.Error("select monthly trns sql err: " + err.Error())
		return
	}
	if count == 0 {
		resp.Trns = []models.Transaction{}
	}
	totalPages := int(math.Ceil(float64(count) / float64(input.Rows)))

	resp.TotalPages = totalPages

	return
}
