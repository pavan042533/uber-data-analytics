package revenue
import(
	"gorm.io/gorm"
)

type Repository interface {
	GetRevenueDistribution(db *gorm.DB) ([]RevenueByPaymentType, error)
}