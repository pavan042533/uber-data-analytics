package revenue

import(
	"uber-data-analytics/models"
	// "uber-data-analytics/public/revenue"

	"gorm.io/gorm"
)
type RevenueRepo struct{}


type RevenueByPaymentType struct {
	PaymentMethod       string    `json:"payment_method"`
	TotalRevenue        float64   `json:"total_revenue"`
    PercentContribution float64    `json:"percent_contribution"`
}

func (r *RevenueRepo) GetRevenueDistribution(db *gorm.DB) ([]RevenueByPaymentType, error){
	var revenueDistribution []RevenueByPaymentType

	err := db.Model(&models.Ride{}).
	            Select("payment_method, SUM(booking_value) as total_revenue, ROUND((SUM(booking_value) / (SELECT SUM(booking_value) FROM rides) * 100),2) as percent_contribution").
				Where("payment_method IS NOT NULL AND payment_method != ?", "").
				Group("payment_method").Order("total_revenue DESC").
				Scan(&revenueDistribution).Error
	if err !=nil{
		return nil, err
	}
	return revenueDistribution, nil
}

func NewRevenueRepo() Repository{
	return &RevenueRepo{}
}

