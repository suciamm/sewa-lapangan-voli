-- name: CreateCourt :execresult
INSERT INTO courts (owner_id, name, description, address, city, price_per_hour, status, avg_rating, total_reviews)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetCourtByID :one
SELECT * FROM courts WHERE id = ? LIMIT 1;

-- name: ListCourts :many
SELECT * FROM courts WHERE status = 'active' ORDER BY created_at DESC;

-- name: ListCourtsByOwner :many
SELECT * FROM courts WHERE owner_id = ? ORDER BY created_at DESC;

-- name: ListCourtsByCity :many
SELECT * FROM courts WHERE city = ? AND status = 'active' ORDER BY created_at DESC;

-- name: UpdateCourt :exec
UPDATE courts
SET name = ?, description = ?, address = ?, city = ?, price_per_hour = ?, status = ?
WHERE id = ?;

-- name: UpdateCourtRating :exec
UPDATE courts
SET avg_rating = ?, total_reviews = ?
WHERE id = ?;

-- name: DeleteCourt :exec
DELETE FROM courts WHERE id = ?;
