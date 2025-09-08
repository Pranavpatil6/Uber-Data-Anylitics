package pkg

import (
	"github.com/pranavpatil6/config"
	"github.com/pranavpatil6/models"
)

func GetHighestRatedVehicleTypes() ([]VehicleStruct, error) {
	var results []VehicleStruct

	err := config.DB.Model(&models.Ride{}).
		Select("vehicle_type, AVG(customer_rating) as average_rating, COUNT(*) as total_rides").
		Where("customer_rating IS NOT NULL AND vehicle_type IS NOT NULL").
		Group("vehicle_type").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}
	return results, nil
}
