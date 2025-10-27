package revenue
import(
	"uber-data-analytics/config"
	"uber-data-analytics/pkg"

	"github.com/gofiber/fiber/v2"
)

func RevenueDistribution(c *fiber.Ctx) error{
	revenueSplit, err:= pkg.RevenueRepo.GetRevenueDistribution(config.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	highestGrossingPaymentMethod := revenueSplit[0]

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"revenue_distribution": revenueSplit,
		"highest_grossing_payment_method": highestGrossingPaymentMethod,
	})
}