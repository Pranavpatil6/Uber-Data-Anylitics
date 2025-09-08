package routes

import (
	"github.com/HOA-Share/hoashare-go-common/http_server"
)

func MountRoutes(app *http_server.Router) {
	api := app.Group("/api")


	api.Get("/most_payment_type", PaymentHandler())


	cancel_api := api.Group("/cancellations")
	cancel_api.Get("/customer", CustomerCancellationHandler())	
	cancel_api.Get("/driver", DriverCancellationHandler())


	ratings_api:= api.Group("/ratings")
	ratings_api.Get("/summary", RatingsSummaryHandler())
	ratings_api.Get("/vehicle-type", VehicleTypeRatingsHandler())


	queries := app.Group("/analytics")
	
	queries.Get("/cancellations/customer-reasons", GetCustomerCancellationReasonsHandler())
	queries.Get("/cancellations/driver-reasons", GetDriverCancellationReasonsHandler())


	queries.Get("/ratings/overall-customer", GetOverallCustomerRatingHandler())
	queries.Get("/ratings/customers-by-rating", GetCustomerRatingsByRiderHandler())
	queries.Get("/ratings/overall-driver", GetOverallDriverRatingHandler())
	queries.Get("/ratings/highest-vehicle-types", GetHighestRatedVehicleTypesHandler())

}
