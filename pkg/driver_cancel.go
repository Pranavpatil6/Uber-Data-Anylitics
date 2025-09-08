package pkg

import (
	"github.com/pranavpatil6/config"
	"github.com/pranavpatil6/models"
)

func GetDriverCancellationReasons() ([]CancellationReason, error) {
	var results []CancellationReason

	err := config.DB.Model(&models.Ride{}).
		Select("driver_cancel_reason as reason, COUNT(*) as count").
		Where("cancelled_by_driver = ? AND driver_cancel_reason IS NOT NULL AND driver_cancel_reason != ?", 1, "").
		Group("driver_cancel_reason").
		Order("count DESC").
		Scan(&results).Error

	if err != nil {
		return nil,err
	}
	return results, nil
}
