package cancellations

import (
	// "uber-data-analytics/public/cancellations"

	"gorm.io/gorm"
)

type Repository interface {
	GetCancilationsByCustomers(db *gorm.DB) ([]CustomerCancellations, error)
	GetCancellationsByDrivers(db *gorm.DB) ([]DriverCancellations, error)
	GetFrequentCancellations(db *gorm.DB) ([]FrequentCancellations, error)
}