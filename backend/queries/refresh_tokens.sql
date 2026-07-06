-- name: CreateRefreshToken :execresult
INSERT INTO refresh_tokens (user_id, token, expires_at)
VALUES (?, ?, ?);

-- name: GetRefreshToken :one
SELECT * FROM refresh_tokens WHERE token = ? LIMIT 1;

-- name: DeleteRefreshToken :exec
DELETE FROM refresh_tokens WHERE token = ?;

-- name: DeleteAllUserTokens :exec
DELETE FROM refresh_tokens WHERE user_id = ?;
