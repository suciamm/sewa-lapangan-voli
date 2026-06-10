-- name: CreateRefreshToken :exec
INSERT INTO refresh_tokens (user_id, token, expires_at)
VALUES (?, ?, ?);

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens
WHERE token = ?
  AND expires_at > NOW()
LIMIT 1;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens WHERE token = ?;

-- name: DeleteAllRefreshTokensByUser :exec
DELETE FROM refresh_tokens WHERE user_id = ?;
