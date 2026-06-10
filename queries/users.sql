-- name: GetUserByID :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ? LIMIT 1;

-- name: CreateUser :execresult
INSERT INTO users (name, email, password, phone, role, status)
VALUES (?, ?, ?, ?, ?, ?);

-- name: UpdateUserStatus :exec
UPDATE users
SET status = ?
WHERE id = ?;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = ?
WHERE id = ?;
