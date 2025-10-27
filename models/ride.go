package models

type Ride struct {
	BookingID                     string  `json:"booking_id"`
	Date                          string  `json:"date"`
	Time                          string  `json:"time"`
	BookingStatus                 string  `json:"booking_status"`
	CustomerID                    string  `json:"customer_id"`
	VeichelType                   string  `json:"vehicle_type"`
	PickupLocation                string  `json:"pickup_location"`
	DropLocation                  string  `json:"drop_location"`
	CancelledRidesByCustomer      int  `json:"cancelled_rides_by_customer"`
	ReasonForCancellingByCustomer string  `json:"reason_for_cancelling_by_customer"`
	CancelledRidesByDriver        int  `json:"cancelled_rides_by_driver"`
	DriverCancellationReason      string  `json:"driver_cancellation_reason"`
	IncompleteRides               int     `json:"incomplete_rides"`
	IncompleteRidesReason         string  `json:"incomplete_rides_reason"`
	BookingValue                  int     `json:"booking_value"`
	RideDistance                  float64 `json:"ride_distance"`
	DriverRatings                 float64 `json:"driver_ratings"`
	CustomerRating                float64 `json:"customer_rating"`
	PaymentMethod                 string  `json:"payment_method"`
}