package pkg

type CancellationReason struct {
	Reason     string  `json:"reason"`
	Count      int64   `json:"count"`
	Percentage float64 `json:"percentage"`
}

type VehicleStruct struct {
	VehicleType   string  `json:"vehicle_type"`
	AverageRating float64 `json:"average_rating"`
	TotalRides    int64   `json:"total_rides"`
}
