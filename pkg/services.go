package pkg

import (
	"uber-data-analytics/pkg/cancellations"
	"uber-data-analytics/pkg/ratings"
	"uber-data-analytics/pkg/revenue"
)

var (
	CancellationRepo cancellations.Repository
	RatingRepo ratings.Repository
	RevenueRepo revenue.Repository
)

func Init(){
	CancellationRepo= cancellations.NewCancellationRepo()
	RatingRepo= ratings.NewRatingRepo()
	RevenueRepo= revenue.NewRevenueRepo()
}