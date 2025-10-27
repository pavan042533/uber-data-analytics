package ratings
import (
	"uber-data-analytics/config"
	"uber-data-analytics/pkg"

	"github.com/gofiber/fiber/v2"
)

func RatingsSummary(c *fiber.Ctx) error {
	avgCustomerRating,err := pkg.RatingRepo.GetAverageCustomerRating(config.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	avrRatingPerCustomer,err := pkg.RatingRepo.GetAverageRatingPerCustomer(config.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	avgDriverRating,err := pkg.RatingRepo.GetAverageDriverRating(config.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Average Customer Rating":        avgCustomerRating,
		"Average Rating Per Customer":    avrRatingPerCustomer,
		"Average Driver Rating":          avgDriverRating,
	})
}

func RatingsByVeichelType(c *fiber.Ctx) error{
	HighestRatedVehicleByCustomer, err:= pkg.RatingRepo.GetHighestRatedVehicleByCustomer(config.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	SortedVeichelTpes, err:= pkg.RatingRepo.GetSortedVehicleTypesByRating(config.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	MostSatisfiedVeichelTypeByDriverws, err:= pkg.RatingRepo.GetMostSatisfiedVehicleTypeByDrivers(config.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Highest Rated Vehicle By Customer": HighestRatedVehicleByCustomer,
		"Sorted Vehicle Types By Rating": SortedVeichelTpes,
		"Most Satisfied Vehicle Type By Drivers": MostSatisfiedVeichelTypeByDriverws,
	})
}