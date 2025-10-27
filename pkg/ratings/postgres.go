package ratings

import (
	"uber-data-analytics/models"
	// "uber-data-analytics/public/cancellations"

	"gorm.io/gorm"
)

type RatingRepo struct{}

func (r *RatingRepo) GetAverageCustomerRating(db *gorm.DB) (float64, error) {
	var avgRating float64
	err := db.Model(&models.Ride{}).Select("AVG(customer_rating)").Where("customer_rating IS NOT NULL").Scan(&avgRating).Error
	if err != nil {
		return 0, err
	}
	return avgRating, nil
}

func (r *RatingRepo) GetAverageRatingPerCustomer(db *gorm.DB) (map[string]float64, error) {
	type Result struct {
		CustomerID string
		AvgRating  float64
	}
	var results []Result
	err := db.Model(&models.Ride{}).Select("customer_id, AVG(customer_rating) as avg_rating").Where("customer_rating IS NOT NULL").Group("customer_id").Scan(&results).Error
	if err != nil {
		return nil, err
	}
	avgRatings := make(map[string]float64)
	for _, res := range results {
		avgRatings[res.CustomerID] = res.AvgRating
	}
	return avgRatings, nil
}

func (r *RatingRepo) GetAverageDriverRating(db *gorm.DB) (float64, error) {
	var avgRating float64
	err := db.Model(&models.Ride{}).Select("AVG(driver_ratings)").Where("driver_ratings IS NOT NULL").Scan(&avgRating).Error
	if err != nil {
		return 0, err
	}
	return avgRating, nil
}

type VehicleRating struct {
	VehicleType string  `json:"vehicle_type"`
	TotalRatings int64  `json:"total_ratings"`
	AvgRating   float64 `json:"avg_rating"`
}

func (r *RatingRepo) GetHighestRatedVehicleByCustomer(db *gorm.DB) (VehicleRating, error) {
	var vehicleRating VehicleRating
	err := db.Model(&models.Ride{}).
	          Select("vehicle_type, COUNT(customer_rating) as total_ratings, AVG(customer_rating) as avg_rating").
			  Where("customer_rating IS NOT NULL").
			  Group("vehicle_type").
			  Order("avg_rating DESC").
			  Limit(1).
			  Scan(&vehicleRating).Error
	if err != nil {
		return VehicleRating{}, err
	}
	return vehicleRating, nil
}

func (r *RatingRepo) GetSortedVehicleTypesByRating(db *gorm.DB) ([]VehicleRating, error) {
	var vehicleRatings []VehicleRating
	err := db.Model(&models.Ride{}).
	          Select("vehicle_type, COUNT(customer_rating) as total_ratings, AVG(customer_rating) as avg_rating").
			  Where("customer_rating IS NOT NULL").
			  Group("vehicle_type").
			  Order("avg_rating DESC , total_ratings DESC ").
			  Scan(&vehicleRatings).Error
	if err != nil {
		return nil, err
	}
	return vehicleRatings, nil
}

func (r *RatingRepo) GetMostSatisfiedVehicleTypeByDrivers(db *gorm.DB) (VehicleRating, error) {
	var vehicleRating VehicleRating
	err := db.Model(&models.Ride{}).
	          Select("vehicle_type, COUNT(driver_ratings) as total_ratings, AVG(driver_ratings) as avg_rating").
			  Where("driver_ratings IS NOT NULL").
			  Group("vehicle_type").
			  Order("avg_rating DESC").
			  Limit(1).
			  Scan(&vehicleRating).Error
	if err != nil {
		return VehicleRating{}, err
	}
	return vehicleRating, nil
}

func NewRatingRepo() Repository {
	return &RatingRepo{}
}