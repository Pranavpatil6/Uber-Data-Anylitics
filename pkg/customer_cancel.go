package pkg

import (
	"github.com/pranavpatil6/config"
	"github.com/pranavpatil6/models"
)


func GetCustomerCancellationReasons() ([]CancellationReason,error){
	var results []CancellationReason

	err := config.DB.Model(&models.Ride{}).
		Select("customer_cancel_reason as reason, COUNT(*) as count").
		Where("cancelled_by_customer = ? AND customer_cancel_reason IS NOT NULL AND customer_cancel_reason != ?", 1, "").
		Group("customer_cancel_reason").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}
	return results, nil
}
