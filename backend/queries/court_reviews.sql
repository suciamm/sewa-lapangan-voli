-- name: CreateCourtReview :execresult
INSERT INTO court_reviews (booking_id, court_id, penyewa_id, rating, comment)
VALUES (?, ?, ?, ?, ?);

-- name: GetReviewByID :one
SELECT * FROM court_reviews WHERE id = ? LIMIT 1;

-- name: GetReviewByBookingID :one
SELECT * FROM court_reviews WHERE booking_id = ? LIMIT 1;

-- name: ListReviewsByCourt :many
SELECT * FROM court_reviews WHERE court_id = ? ORDER BY created_at DESC;

-- name: ListReviewsByPenyewa :many
SELECT * FROM court_reviews WHERE penyewa_id = ? ORDER BY created_at DESC;
