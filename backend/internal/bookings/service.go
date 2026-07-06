package bookings

import (
	"context"
	"errors"
	"time"

	sqlc_db "sewa-lapangan-voli/db"
)

func CreateBookingService(ctx context.Context, q *sqlc_db.Queries, penyewaID int64, req CreateBookingRequest) (BookingResponse, error) {
	court, err := q.GetCourtByID(ctx, req.CourtID)
	if err != nil {
		return BookingResponse{}, errors.New("lapangan tidak ditemukan")
	}

	if court.Status != "active" {
		return BookingResponse{}, errors.New("lapangan tidak tersedia")
	}

	bookingDate, err := time.Parse("2006-01-02", req.BookingDate)
	if err != nil {
		return BookingResponse{}, errors.New("format tanggal tidak valid (YYYY-MM-DD)")
	}

	startTime, err := time.Parse("15:04", req.StartTime)
	if err != nil {
		return BookingResponse{}, errors.New("format waktu mulai tidak valid (HH:MM)")
	}

	endTime, err := time.Parse("15:04", req.EndTime)
	if err != nil {
		return BookingResponse{}, errors.New("format waktu selesai tidak valid (HH:MM)")
	}

	duration := int32(endTime.Sub(startTime).Hours())
	if duration <= 0 {
		return BookingResponse{}, errors.New("waktu selesai harus setelah waktu mulai")
	}

	totalPrice := float64(duration) * court.PricePerHour

	conflicts, err := q.CheckCourtAvailability(ctx, sqlc_db.CheckCourtAvailabilityParams{
		CourtID:     req.CourtID,
		BookingDate: bookingDate,
		StartTime:   startTime,
		EndTime:     endTime,
	})
	if err == nil && len(conflicts) > 0 {
		return BookingResponse{}, errors.New("jadwal sudah terbooking")
	}

	result, err := q.CreateBooking(ctx, sqlc_db.CreateBookingParams{
		CourtID:       req.CourtID,
		PenyewaID:     penyewaID,
		BookingDate:   bookingDate,
		StartTime:     startTime,
		EndTime:       endTime,
		DurationHours: duration,
		TotalPrice:    totalPrice,
		Status:        sqlc_db.BookingsStatus("pending_payment"),
	})
	if err != nil {
		return BookingResponse{}, errors.New("gagal membuat booking")
	}

	bookingID, err := result.LastInsertId()
	if err != nil {
		return BookingResponse{}, errors.New("gagal mendapatkan ID booking")
	}

	return BookingResponse{
		ID:            bookingID,
		CourtID:       req.CourtID,
		PenyewaID:     penyewaID,
		BookingDate:   req.BookingDate,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		DurationHours: duration,
		TotalPrice:    totalPrice,
		Status:        "pending_payment",
		CreatedAt:     time.Now(),
		CourtName:     court.Name,
	}, nil
}

func GetBookingService(ctx context.Context, q *sqlc_db.Queries, id int64, userID int64, role string) (BookingResponse, error) {
	booking, err := q.GetBookingByID(ctx, id)
	if err != nil {
		return BookingResponse{}, errors.New("booking tidak ditemukan")
	}

	if role == "penyewa" && booking.PenyewaID != userID {
		return BookingResponse{}, errors.New("tidak diizinkan")
	}

	if role == "owner" {
		court, _ := q.GetCourtByID(ctx, booking.CourtID)
		if court.OwnerID != userID {
			return BookingResponse{}, errors.New("tidak diizinkan")
		}
	}

	court, _ := q.GetCourtByID(ctx, booking.CourtID)

	return BookingResponse{
		ID:            booking.ID,
		CourtID:       booking.CourtID,
		PenyewaID:     booking.PenyewaID,
		BookingDate:   booking.BookingDate.Format("2006-01-02"),
		StartTime:     booking.StartTime.Format("15:04"),
		EndTime:       booking.EndTime.Format("15:04"),
		DurationHours: booking.DurationHours,
		TotalPrice:    booking.TotalPrice,
		Status:        string(booking.Status),
		CreatedAt:     booking.CreatedAt,
		CourtName:     court.Name,
	}, nil
}

func ListPenyewaBookingsService(ctx context.Context, q *sqlc_db.Queries, penyewaID int64) ([]BookingResponse, error) {
	bookings, err := q.ListBookingsByPenyewa(ctx, penyewaID)
	if err != nil {
		return nil, errors.New("gagal mendapatkan daftar booking")
	}

	resp := make([]BookingResponse, 0, len(bookings))
	for _, b := range bookings {
		court, _ := q.GetCourtByID(ctx, b.CourtID)
		resp = append(resp, BookingResponse{
			ID:            b.ID,
			CourtID:       b.CourtID,
			PenyewaID:     b.PenyewaID,
			BookingDate:   b.BookingDate.Format("2006-01-02"),
			StartTime:     b.StartTime.Format("15:04"),
			EndTime:       b.EndTime.Format("15:04"),
			DurationHours: b.DurationHours,
			TotalPrice:    b.TotalPrice,
			Status:        string(b.Status),
			CreatedAt:     b.CreatedAt,
			CourtName:     court.Name,
		})
	}
	return resp, nil
}

func ListOwnerBookingsService(ctx context.Context, q *sqlc_db.Queries, ownerID int64) ([]BookingResponse, error) {
	bookings, err := q.ListBookingsByOwner(ctx, ownerID)
	if err != nil {
		return nil, errors.New("gagal mendapatkan daftar booking")
	}

	resp := make([]BookingResponse, 0, len(bookings))
	for _, b := range bookings {
		court, _ := q.GetCourtByID(ctx, b.CourtID)
		resp = append(resp, BookingResponse{
			ID:            b.ID,
			CourtID:       b.CourtID,
			PenyewaID:     b.PenyewaID,
			BookingDate:   b.BookingDate.Format("2006-01-02"),
			StartTime:     b.StartTime.Format("15:04"),
			EndTime:       b.EndTime.Format("15:04"),
			DurationHours: b.DurationHours,
			TotalPrice:    b.TotalPrice,
			Status:        string(b.Status),
			CreatedAt:     b.CreatedAt,
			CourtName:     court.Name,
		})
	}
	return resp, nil
}

func UpdateBookingStatusService(ctx context.Context, q *sqlc_db.Queries, id int64, status string, userID int64, role string) error {
	booking, err := q.GetBookingByID(ctx, id)
	if err != nil {
		return errors.New("booking tidak ditemukan")
	}

	if role == "penyewa" {
		if status != "cancelled" {
			return errors.New("hanya bisa cancel booking")
		}
		if booking.PenyewaID != userID {
			return errors.New("tidak diizinkan")
		}
	} else if role == "owner" {
		court, _ := q.GetCourtByID(ctx, booking.CourtID)
		if court.OwnerID != userID {
			return errors.New("tidak diizinkan")
		}
	}

	return q.UpdateBookingStatus(ctx, sqlc_db.UpdateBookingStatusParams{
		ID:     id,
		Status: sqlc_db.BookingsStatus(status),
	})
}
