package courts

import "time"

type CreateCourtRequest struct {
	Name         string  `json:"name"          binding:"required"`
	Description  string  `json:"description"`
	Address      string  `json:"address"       binding:"required"`
	City         string  `json:"city"          binding:"required"`
	PricePerHour float64 `json:"price_per_hour" binding:"required,min=0"`
}

type UpdateCourtRequest struct {
	Name         string  `json:"name"          binding:"required"`
	Description  string  `json:"description"`
	Address      string  `json:"address"       binding:"required"`
	City         string  `json:"city"          binding:"required"`
	PricePerHour float64 `json:"price_per_hour" binding:"required,min=0"`
	Status       string  `json:"status"        binding:"required,oneof=active inactive maintenance"`
}

type CourtImageRequest struct {
	ImageURL  string `json:"image_url"  binding:"required"`
	IsPrimary bool   `json:"is_primary"`
}

type CourtResponse struct {
	ID           int64            `json:"id"`
	OwnerID      int64            `json:"owner_id"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Address      string           `json:"address"`
	City         string           `json:"city"`
	PricePerHour float64          `json:"price_per_hour"`
	Status       string           `json:"status"`
	AvgRating    *float64         `json:"avg_rating"`
	TotalReviews int32            `json:"total_reviews"`
	CreatedAt    time.Time        `json:"created_at"`
	Images       []CourtImageResp `json:"images"`
}

type CourtImageResp struct {
	ID        int64  `json:"id"`
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}
