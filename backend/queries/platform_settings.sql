-- name: GetPlatformSettings :one
SELECT * FROM platform_settings ORDER BY id DESC LIMIT 1;

-- name: UpdatePlatformSettings :exec
UPDATE platform_settings SET fee_percent = ? WHERE id = ?;
