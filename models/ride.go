package models

type Ride struct {
	Date                  string  `json:"date"`
	Time                  string  `json:"time"`
	BookingID             string  `json:"booking_id"`
	BookingStatus         string  `json:"booking_status"`
	CustomerID            string  `json:"customer_id"`
	VehicleType           string  `json:"vehicle_type"`
	PickupLocation        string  `json:"pickup_location"`
	DropLocation          string  `json:"drop_location"`
	AvgVTAT               float64 `json:"avg_vtat"`
	AvgCTAT               float64 `json:"avg_ctat"`
	CancelledByCustomer   int     `json:"cancelled_by_customer"`
	CustomerCancelReason  string  `json:"customer_cancel_reason"`
	CancelledByDriver     int     `json:"cancelled_by_driver"`
	DriverCancelReason    string  `json:"driver_cancel_reason"`
	IncompleteRides       int     `json:"incomplete_rides"`
	IncompleteRidesReason string  `json:"incomplete_rides_reason"`
	BookingValue          float64 `json:"booking_value"`
	RideDistance          float64 `json:"ride_distance"`
	DriverRatings         float64 `json:"driver_ratings"`
	CustomerRating        float64 `json:"customer_rating"`
	PaymentMethod         string  `json:"payment_method"`
}
