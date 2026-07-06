package payments

import "time"

type CreatePaymentRequest struct {
	BookingID int64  `json:"booking_id" binding:"required"`
}

type MidtransWebhookRequest struct {
	OrderID           string  `json:"order_id"`
	TransactionID     string  `json:"transaction_id"`
	TransactionStatus string  `json:"transaction_status"`
	GrossAmount       string  `json:"gross_amount"`
	PaymentType       string  `json:"payment_type"`
}

type PaymentResponse struct {
	ID             int64     `json:"id"`
	BookingID      int64     `json:"booking_id"`
	MidtransOrderID string   `json:"midtrans_order_id"`
	MidtransTxID   string    `json:"midtrans_tx_id,omitempty"`
	Amount         float64   `json:"amount"`
	FeeAmount      float64   `json:"fee_amount"`
	NetToOwner     float64   `json:"net_to_owner"`
	PaymentMethod  string    `json:"payment_method,omitempty"`
	Status         string    `json:"status"`
	PaidAt         *time.Time `json:"paid_at,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
}
