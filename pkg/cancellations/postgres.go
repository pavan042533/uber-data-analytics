package cancellations

import (
	"uber-data-analytics/models"
	// "uber-data-analytics/public/cancellations"

	"gorm.io/gorm"
)

type CustomerCancellations struct {
	Reason string   `json:"reason"`
	Count  int       `json:"count"`
}
type DriverCancellations struct {
	Reason string     `json:"reason"`
	Count  int        `json:"count"`
}

type FrequentCancellations struct {
	CancellationReason string `json:"cancellation_reason"`
	TotalOccurrences   int    `json:"total_occurrences"`
}

type cancellationRepo struct{}

func (c *cancellationRepo) GetCancilationsByCustomers(db *gorm.DB) ([]CustomerCancellations, error){
	var customerCancellations []CustomerCancellations

	err := db.Model(&models.Ride{}).
	            Select("reason_for_cancelling_by_customer as reason, COUNT(*) as count").
				Where("cancelled_rides_by_customer = ? AND reason_for_cancelling_by_customer IS NOT NULL AND reason_for_cancelling_by_customer != ?", 1, "").
				Group("reason_for_cancelling_by_customer").Order("count DESC").
				Scan(&customerCancellations).Error
	if err !=nil{
		return nil, err
	} else{
		return customerCancellations, nil
	}
}

func (c *cancellationRepo) GetCancellationsByDrivers(db *gorm.DB) ([]DriverCancellations, error){
	var driverCancellations []DriverCancellations
	err := db.Model(&models.Ride{}).
	            Select("driver_cancellation_reason as reason, COUNT(*) as count").
				Where("cancelled_rides_by_driver = ? AND driver_cancellation_reason IS NOT NULL AND driver_cancellation_reason != ?", 1, "").
				Group("driver_cancellation_reason").Order("count DESC").
				Scan(&driverCancellations).Error
	if err !=nil{
		return nil, err
	} else{
		return driverCancellations, nil
	}
}


func (c *cancellationRepo) GetFrequentCancellations(db *gorm.DB) ([]FrequentCancellations, error){
	var frequentCancellations []FrequentCancellations
	err := db.Raw(`
    SELECT 
        reason AS cancellation_reason,
        SUM(total) AS total_occurrences
    FROM (
        SELECT 
            reason_for_cancelling_by_customer AS reason,
            COUNT(*) AS total
        FROM rides
        WHERE cancelled_rides_by_customer = 1
        GROUP BY reason_for_cancelling_by_customer

        UNION ALL

        SELECT 
            driver_cancellation_reason AS reason,
            COUNT(*) AS total
        FROM rides
        WHERE cancelled_rides_by_driver=1
        GROUP BY driver_cancellation_reason
    ) AS combined_reasons
    GROUP BY reason
    ORDER BY total_occurrences DESC
    LIMIT 3
`).Scan(&frequentCancellations).Error
	if err !=nil{
		return nil, err
	} else{
		return frequentCancellations, nil
	}
}

func NewCancellationRepo() Repository{
	return &cancellationRepo{}
}
