package pkg

import (
	"github.com/pranavpatil6/config"
	"github.com/pranavpatil6/models"
)

// func GetOverallCustomerRating() (float64, error) {
// 	var avgcustomerRating float64

// 	config.DB.Model(&models.Ride{}).
// 		Select("AVG(customer_rating)").
// 		Scan(&avgcustomerRating).Error

// 	var avgRating float64

// 	config.DB.Model(&models.Ride{}).
// 		Select("AVG(customer_rating)").
// 		Scan(&avgRating).Error

// 	return avgcustomerRating, nil
// }

func Ratings_summary() (map[string]float64,error) {
	ratings_summary := make(map[string]float64)
	var avgDriverRating float64

	config.DB.Model(&models.Ride{}).
		Select("AVG(driver_ratings)").
		Scan(&avgDriverRating)
	var avgCustomerRating float64

	config.DB.Model(&models.Ride{}).
		Select("AVG(customer_rating)").
		Scan(&avgCustomerRating)

	ratings_summary["average driver rating"] = float64(avgDriverRating)
	ratings_summary["average customer rating"] = float64(avgCustomerRating)
	return ratings_summary,nil
}
