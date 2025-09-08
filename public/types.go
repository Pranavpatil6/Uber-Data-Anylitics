package routes

type MostPaymentTypeQuerry struct {
	PaymentType string `query:"payment_type "`
}

type CancellationReason struct {
	Reason     string  `json:"reason"`
	Count      int64   `json:"count"`
	Percentage float64 `json:"percentage"`
}
