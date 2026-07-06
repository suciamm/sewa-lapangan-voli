-- name: CreateBooking :execresult
INSERT INTO bookings (court_id, penyewa_id, booking_date, start_time, end_time, duration_hours, total_price, status)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetBookingByID :one
SELECT * FROM bookings WHERE id = ? LIMIT 1;

-- name: ListBookingsByPenyewa :many
SELECT * FROM bookings WHERE penyewa_id = ? ORDER BY created_at DESC;

-- name: ListBookingsByCourt :many
SELECT * FROM bookings WHERE court_id = ? ORDER BY created_at DESC;

-- name: ListBookingsByOwner :many
SELECT b.* FROM bookings b
JOIN courts c ON b.court_id = c.id
WHERE c.owner_id = ?
ORDER BY b.created_at DESC;

-- name: CheckCourtAvailability :many
SELECT * FROM bookings
WHERE court_id = ?
AND booking_date = ?
AND status NOT IN ('cancelled')
AND (
    (start_time < ? AND end_time > ?) OR
    (start_time < ? AND end_time > ?) OR
    (start_time >= ? AND end_time <= ?)
);

-- name: UpdateBookingStatus :exec
UPDATE bookings SET status = ? WHERE id = ?;

-- name: DeleteBooking :exec
DELETE FROM bookings WHERE id = ?;
