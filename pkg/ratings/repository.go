package ratings
import (
	"gorm.io/gorm"
)

type Repository interface {
	GetAverageCustomerRating(db *gorm.DB) (float64, error)
	GetAverageRatingPerCustomer(db *gorm.DB) (map[string]float64, error)
	GetAverageDriverRating(db *gorm.DB) (float64, error)
	GetHighestRatedVehicleByCustomer(db *gorm.DB) (VehicleRating, error)
	GetSortedVehicleTypesByRating(db *gorm.DB) ([]VehicleRating, error)
	GetMostSatisfiedVehicleTypeByDrivers(db *gorm.DB) (VehicleRating, error)
}

