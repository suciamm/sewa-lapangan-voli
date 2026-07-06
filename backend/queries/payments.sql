-- name: CreatePayment :execresult
INSERT INTO payments (booking_id, midtrans_order_id, amount, fee_amount, net_to_owner, payment_method, status)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetPaymentByID :one
SELECT * FROM payments WHERE id = ? LIMIT 1;

-- name: GetPaymentByBookingID :one
SELECT * FROM payments WHERE booking_id = ? LIMIT 1;

-- name: GetPaymentByMidtransOrderID :one
SELECT * FROM payments WHERE midtrans_order_id = ? LIMIT 1;

-- name: UpdatePayment :exec
UPDATE payments
SET midtrans_tx_id = ?, payment_method = ?, status = ?, paid_at = ?
WHERE id = ?;

-- name: ListPayments :many
SELECT * FROM payments ORDER BY created_at DESC;
