package repository

import (
	"github.com/pranavpatil6/config"
	"github.com/pranavpatil6/models"
)

// -----------------
type CancellationReason struct {
	Reason     string  `json:"reason"`
	Count      int64   `json:"count"`
	Percentage float64 `json:"percentage"`
}

func GetCustomerCancellationReasons() ([]CancellationReason, error) {
	var results []CancellationReason

	err := config.DB.Model(&models.Ride{}).
		Select("customer_cancel_reason as reason, COUNT(*) as count, (COUNT(*) * 100.0 / (SELECT COUNT(*) FROM rides WHERE cancelled_by_customer = 1)) as percentage").
		Where("cancelled_by_customer = ? AND customer_cancel_reason IS NOT NULL AND customer_cancel_reason != ?", 1, "").
		Group("customer_cancel_reason").
		Order("count DESC").
		Scan(&results).Error

	return results, err
}

func GetDriverCancellationReasons() ([]CancellationReason, error) {
	var results []CancellationReason

	err := config.DB.Model(&models.Ride{}).
		Select("driver_cancel_reason as reason, COUNT(*) as count, (COUNT(*) * 100.0 / (SELECT COUNT(*) FROM rides WHERE cancelled_by_driver = 1)) as percentage").
		Where("cancelled_by_driver = ? AND driver_cancel_reason IS NOT NULL AND driver_cancel_reason != ?", 1, "").
		Group("driver_cancel_reason").
		Order("count DESC").
		Scan(&results).Error

	return results, err
}

func GetOverallCustomerRating() ( float64, error) {
	var avgRating float64

	err := config.DB.Model(&models.Ride{}).
		Select("AVG(customer_rating)").
		Where("customer_rating IS NOT NULL").
		Scan(&avgRating).Error

	return avgRating, err
}


type CustomerRatingStats struct {
	CustomerID    string  `json:"customer_id"`
	AverageRating float64 `json:"average_rating"`
	TotalRides    int64   `json:"total_rides"`
	TotalRatings  int64   `json:"total_ratings"`
}

func GetCustomerRatingsByRider() ([]CustomerRatingStats, error) {
	var results []CustomerRatingStats

	err := config.DB.Model(&models.Ride{}).
		Select("customer_id, AVG(customer_rating) as average_rating, COUNT(*) as total_rides, COUNT(customer_rating) as total_ratings").
		Where("customer_rating IS NOT NULL").
		Group("customer_id").
		Having("COUNT(customer_rating) >= ?", 1). 
		Order("average_rating DESC").
		Scan(&results).Error

	return results, err
}


func GetOverallDriverRating() (string, error) {
	var avgRating string

	err := config.DB.Model(&models.Ride{}).
		Select("AVG(driver_ratings)").
		Where("driver_ratings IS NOT NULL").
		Scan(&avgRating).Error

	return avgRating, err
}


type VehicleRatingStats struct {
	VehicleType   string  `json:"vehicle_type"`
	AverageRating float64 `json:"average_rating"`
	TotalRides    int64   `json:"total_rides"`
	TotalRatings  int64   `json:"total_ratings"`
}

func GetHighestRatedVehicleTypes() ([]VehicleRatingStats, error) {
	var results []VehicleRatingStats

	err := config.DB.Model(&models.Ride{}).
		Select("vehicle_type, AVG(customer_rating) as average_rating, COUNT(*) as total_rides, COUNT(customer_rating) as total_ratings").
		Where("customer_rating IS NOT NULL AND vehicle_type IS NOT NULL").
		Group("vehicle_type").
		Having("COUNT(customer_rating) >= ?", 1).
		Order("average_rating DESC").
		Scan(&results).Error

	return results, err
}
