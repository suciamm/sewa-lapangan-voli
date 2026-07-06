package reviews

import (
	"context"
	"database/sql"
	"errors"
	"time"

	sqlc_db "sewa-lapangan-voli/db"
)

func CreateReviewService(ctx context.Context, q *sqlc_db.Queries, penyewaID int64, req CreateReviewRequest) (ReviewResponse, error) {
	booking, err := q.GetBookingByID(ctx, req.BookingID)
	if err != nil {
		return ReviewResponse{}, errors.New("booking tidak ditemukan")
	}

	if booking.PenyewaID != penyewaID {
		return ReviewResponse{}, errors.New("tidak diizinkan")
	}

	if booking.Status != "completed" {
		return ReviewResponse{}, errors.New("hanya bisa review booking yang sudah selesai")
	}

	existing, _ := q.GetReviewByBookingID(ctx, req.BookingID)
	if existing.ID > 0 {
		return ReviewResponse{}, errors.New("sudah ada review untuk booking ini")
	}

	result, err := q.CreateCourtReview(ctx, sqlc_db.CreateCourtReviewParams{
		BookingID: req.BookingID,
		CourtID:   booking.CourtID,
		PenyewaID: penyewaID,
		Rating:    req.Rating,
		Comment:   toNullString(req.Comment),
	})
	if err != nil {
		return ReviewResponse{}, errors.New("gagal membuat review")
	}

	reviewID, err := result.LastInsertId()
	if err != nil {
		return ReviewResponse{}, errors.New("gagal mendapatkan ID review")
	}

	reviews, err := q.ListReviewsByCourt(ctx, booking.CourtID)
	if err == nil {
		total := int32(len(reviews))
		sum := 0.0
		for _, r := range reviews {
			sum += float64(r.Rating)
		}
		avg := sum / float64(total)
		_ = q.UpdateCourtRating(ctx, sqlc_db.UpdateCourtRatingParams{
			ID:           booking.CourtID,
			AvgRating:    avg,
			TotalReviews: total,
		})
	}

	user, _ := q.GetUserByID(ctx, penyewaID)

	return ReviewResponse{
		ID:          reviewID,
		BookingID:   req.BookingID,
		CourtID:     booking.CourtID,
		PenyewaID:   penyewaID,
		Rating:      req.Rating,
		Comment:     req.Comment,
		CreatedAt:   time.Now(),
		PenyewaName: user.Name,
	}, nil
}

func ListCourtReviewsService(ctx context.Context, q *sqlc_db.Queries, courtID int64) ([]ReviewResponse, error) {
	reviews, err := q.ListReviewsByCourt(ctx, courtID)
	if err != nil {
		return nil, errors.New("gagal mendapatkan daftar review")
	}

	resp := make([]ReviewResponse, 0, len(reviews))
	for _, r := range reviews {
		user, _ := q.GetUserByID(ctx, r.PenyewaID)
		resp = append(resp, ReviewResponse{
			ID:          r.ID,
			BookingID:   r.BookingID,
			CourtID:     r.CourtID,
			PenyewaID:   r.PenyewaID,
			Rating:      r.Rating,
			Comment:     nullStringToString(r.Comment),
			CreatedAt:   r.CreatedAt,
			PenyewaName: user.Name,
		})
	}
	return resp, nil
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func nullStringToString(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}
