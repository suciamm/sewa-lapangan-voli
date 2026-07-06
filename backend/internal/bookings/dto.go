package bookings

import "time"

type CreateBookingRequest struct {
	CourtID     int64  `json:"court_id"      binding:"required"`
	BookingDate string `json:"booking_date"  binding:"required"`
	StartTime   string `json:"start_time"    binding:"required"`
	EndTime     string `json:"end_time"      binding:"required"`
}

type UpdateBookingStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=paid active completed cancelled"`
}

type BookingResponse struct {
	ID            int64     `json:"id"`
	CourtID       int64     `json:"court_id"`
	PenyewaID     int64     `json:"penyewa_id"`
	BookingDate   string    `json:"booking_date"`
	StartTime     string    `json:"start_time"`
	EndTime       string    `json:"end_time"`
	DurationHours int32     `json:"duration_hours"`
	TotalPrice    float64   `json:"total_price"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
	CourtName     string    `json:"court_name"`
}
