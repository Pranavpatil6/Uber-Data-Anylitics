package routes

import (
	"github.com/HOA-Share/hoashare-go-common/http_server"
	"github.com/gofiber/fiber/v2"
	"github.com/pranavpatil6/pkg"
)

func RatingsSummaryHandler() fiber.Handler {
	return http_server.NewHandler[any, any, any]().
		SetRouteAccess(http_server.Public).
		SetHandler(func(data http_server.RequestData[any, any, any]) (any, error) {
			customerRating, err := pkg.Ratings_summary()
			if err != nil {
				return nil, err
			}

			return customerRating, nil
		}).Build()
}
