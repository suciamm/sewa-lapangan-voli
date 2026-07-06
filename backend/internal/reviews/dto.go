package reviews

import "time"

type CreateReviewRequest struct {
	BookingID int64  `json:"booking_id" binding:"required"`
	Rating    int8   `json:"rating"     binding:"required,min=1,max=5"`
	Comment   string `json:"comment"`
}

type ReviewResponse struct {
	ID         int64     `json:"id"`
	BookingID  int64     `json:"booking_id"`
	CourtID    int64     `json:"court_id"`
	PenyewaID  int64     `json:"penyewa_id"`
	Rating     int8      `json:"rating"`
	Comment    string    `json:"comment"`
	CreatedAt  time.Time `json:"created_at"`
	PenyewaName string   `json:"penyewa_name"`
}
