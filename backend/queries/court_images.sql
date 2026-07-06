-- name: CreateCourtImage :execresult
INSERT INTO court_images (court_id, image_url, is_primary)
VALUES (?, ?, ?);

-- name: GetCourtImages :many
SELECT * FROM court_images WHERE court_id = ? ORDER BY is_primary DESC, created_at ASC;

-- name: SetPrimaryImage :exec
UPDATE court_images
SET is_primary = (id = ?)
WHERE court_id = ?;

-- name: DeleteCourtImage :exec
DELETE FROM court_images WHERE id = ?;
