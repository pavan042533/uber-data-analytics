package public

import(
	"github.com/gofiber/fiber/v2"
	"uber-data-analytics/public/cancellations"
	"uber-data-analytics/public/ratings"
	"uber-data-analytics/public/revenue"
)

func MountRoutes(app *fiber.App){
	ApiGroup:= app.Group("/api")
	RevenueApi:=ApiGroup.Group("revenue")
	RevenueApi.Get("/distribution", revenue.RevenueDistribution)

	CancellationRoutes:= ApiGroup.Group("/cancellations")
	CancellationRoutes.Get("/customer", cancellations.CancellationsByCustomer)
	CancellationRoutes.Get("/driver", cancellations.CancellationsByDriver)
	CancellationRoutes.Get("/frequent", cancellations.FrequentCancellations)

	RatingRoutes:=ApiGroup.Group("/ratings")
	RatingRoutes.Get("/summary", ratings.RatingsSummary)
	RatingRoutes.Get("/vehicle-type", ratings.RatingsByVeichelType)
}