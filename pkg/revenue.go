package pkg

import (
	"github.com/pranavpatil6/config"
	"github.com/pranavpatil6/models"
)

func GetMostPaymentType(filterPaymentType string) (any, error) {
	var result struct {
		PaymentMethod string `json:"payment_method"`
		Count         int64  `json:"count"`
	}

	query := config.DB.Model(&models.Ride{}).
		Select("payment_method, COUNT(*) as count").
		Group("payment_method").
		Order("count DESC").
		Limit(1)
	
	if filterPaymentType != "" {
		query = query.Where("payment_method = ?", filterPaymentType)
	}

	err := query.Scan(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
