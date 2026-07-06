package payments

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"time"

	sqlc_db "sewa-lapangan-voli/db"
)

func CreatePaymentService(ctx context.Context, q *sqlc_db.Queries, penyewaID int64, req CreatePaymentRequest) (PaymentResponse, error) {
	booking, err := q.GetBookingByID(ctx, req.BookingID)
	if err != nil {
		return PaymentResponse{}, errors.New("booking tidak ditemukan")
	}

	if booking.PenyewaID != penyewaID {
		return PaymentResponse{}, errors.New("tidak diizinkan")
	}

	if booking.Status != "pending_payment" {
		return PaymentResponse{}, errors.New("booking sudah diproses")
	}

	settings, err := q.GetPlatformSettings(ctx)
	if err != nil {
		return PaymentResponse{}, errors.New("gagal mendapatkan pengaturan platform")
	}

	feeAmount := booking.TotalPrice * (settings.FeePercent / 100)
	netToOwner := booking.TotalPrice - feeAmount

	orderID := fmt.Sprintf("ORD-%d-%d", req.BookingID, rand.Intn(10000))

	result, err := q.CreatePayment(ctx, sqlc_db.CreatePaymentParams{
		BookingID:       req.BookingID,
		MidtransOrderID: orderID,
		Amount:          booking.TotalPrice,
		FeeAmount:       feeAmount,
		NetToOwner:      netToOwner,
		PaymentMethod:   sql.NullString{Valid: false},
		Status:          sqlc_db.PaymentsStatus("pending"),
	})
	if err != nil {
		return PaymentResponse{}, errors.New("gagal membuat payment")
	}

	paymentID, err := result.LastInsertId()
	if err != nil {
		return PaymentResponse{}, errors.New("gagal mendapatkan ID payment")
	}

	return PaymentResponse{
		ID:              paymentID,
		BookingID:       req.BookingID,
		MidtransOrderID: orderID,
		Amount:          booking.TotalPrice,
		FeeAmount:       feeAmount,
		NetToOwner:      netToOwner,
		Status:          "pending",
		CreatedAt:       time.Now(),
	}, nil
}

func ProcessMidtransWebhookService(ctx context.Context, q *sqlc_db.Queries, req MidtransWebhookRequest) error {
	payment, err := q.GetPaymentByMidtransOrderID(ctx, req.OrderID)
	if err != nil {
		return errors.New("payment tidak ditemukan")
	}

	var newStatus sqlc_db.PaymentsStatus
	var paidAt sql.NullTime

	switch req.TransactionStatus {
	case "settlement", "capture":
		newStatus = "settlement"
		paidAt = sql.NullTime{Time: time.Now(), Valid: true}
		if err := q.UpdateBookingStatus(ctx, sqlc_db.UpdateBookingStatusParams{
			ID:     payment.BookingID,
			Status: "paid",
		}); err != nil {
		}
	case "expire":
		newStatus = "expire"
		if err := q.UpdateBookingStatus(ctx, sqlc_db.UpdateBookingStatusParams{
			ID:     payment.BookingID,
			Status: "cancelled",
		}); err != nil {
		}
	case "deny", "cancel":
		newStatus = "failure"
		if err := q.UpdateBookingStatus(ctx, sqlc_db.UpdateBookingStatusParams{
			ID:     payment.BookingID,
			Status: "cancelled",
		}); err != nil {
		}
	default:
		newStatus = payment.Status
	}

	return q.UpdatePayment(ctx, sqlc_db.UpdatePaymentParams{
		ID:            payment.ID,
		MidtransTxID:  toNullString(req.TransactionID),
		PaymentMethod: toNullString(req.PaymentType),
		Status:        newStatus,
		PaidAt:        paidAt,
	})
}

func GetPaymentService(ctx context.Context, q *sqlc_db.Queries, id int64) (PaymentResponse, error) {
	payment, err := q.GetPaymentByID(ctx, id)
	if err != nil {
		return PaymentResponse{}, errors.New("payment tidak ditemukan")
	}

	var paidAt *time.Time
	if payment.PaidAt.Valid {
		paidAt = &payment.PaidAt.Time
	}

	return PaymentResponse{
		ID:              payment.ID,
		BookingID:       payment.BookingID,
		MidtransOrderID: payment.MidtransOrderID,
		MidtransTxID:    nullStringToString(payment.MidtransTxID),
		Amount:          payment.Amount,
		FeeAmount:       payment.FeeAmount,
		NetToOwner:      payment.NetToOwner,
		PaymentMethod:   nullStringToString(payment.PaymentMethod),
		Status:          string(payment.Status),
		PaidAt:          paidAt,
		CreatedAt:       payment.CreatedAt,
	}, nil
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

func nullTimeToPtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
