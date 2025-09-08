package routes

import (
	"github.com/HOA-Share/hoashare-go-common/http_server"
	"github.com/gofiber/fiber/v2"
	"github.com/pranavpatil6/pkg/repository"
)

func GetCustomerCancellationReasonsHandler() fiber.Handler {
	return http_server.NewHandler[any, any, any]().
		SetRouteAccess(http_server.Public).
		SetHandler(func(data http_server.RequestData[any, any, any]) (any, error) {
			return repository.GetCustomerCancellationReasons()
		}).Build()
}

func GetDriverCancellationReasonsHandler() fiber.Handler {
	return http_server.NewHandler[any, any, any]().
		SetRouteAccess(http_server.Public).
		SetHandler(func(data http_server.RequestData[any, any, any]) (any, error) {
			return repository.GetDriverCancellationReasons()
		}).Build()
}

func GetOverallCustomerRatingHandler() fiber.Handler {
	return http_server.NewHandler[any, any, any]().
		SetRouteAccess(http_server.Public).
		SetHandler(func(data http_server.RequestData[any, any, any]) (any, error) {
			rating, err := repository.GetOverallCustomerRating()
			if err != nil {
				return nil, err
			}
			return map[string]interface{}{
				"overall_customer_rating": rating,
			}, nil
		}).Build()
}

func GetCustomerRatingsByRiderHandler() fiber.Handler {
	return http_server.NewHandler[any, any, any]().
		SetRouteAccess(http_server.Public).
		SetHandler(func(data http_server.RequestData[any, any, any]) (any, error) {
			return repository.GetCustomerRatingsByRider()
		}).Build()
}

func GetOverallDriverRatingHandler() fiber.Handler {
	return http_server.NewHandler[any, any, any]().
		SetRouteAccess(http_server.Public).
		SetHandler(func(data http_server.RequestData[any, any, any]) (any, error) {
			rating, err := repository.GetOverallDriverRating()
			if err != nil {
				return nil, err
			}
			return map[string]interface{}{
				"overall_driver_rating": rating,
			}, nil
		}).Build()
}

func GetHighestRatedVehicleTypesHandler() fiber.Handler {
	return http_server.NewHandler[any, any, any]().
		SetRouteAccess(http_server.Public).
		SetHandler(func(data http_server.RequestData[any, any, any]) (any, error) {
			return repository.GetHighestRatedVehicleTypes()
		}).Build()
}
