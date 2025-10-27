package cancellations

import (
	"uber-data-analytics/config"
	"uber-data-analytics/pkg"

	"github.com/gofiber/fiber/v2"
)

func CancellationsByCustomer(c *fiber.Ctx) error {
	result, err := pkg.CancellationRepo.GetCancilationsByCustomers(config.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Breakdown of cancellations by customer reasons": result,
	})
}

func CancellationsByDriver(c *fiber.Ctx) error{
	result, err := pkg.CancellationRepo.GetCancellationsByDrivers(config.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Breakdown of cancellations by driver reasons": result,
	})
}

func FrequentCancellations(c *fiber.Ctx) error {
	result, err:= pkg.CancellationRepo.GetFrequentCancellations(config.DB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"List  top three frequent cancellations": result,
	})
}