-- name: CreateNotification :execresult
INSERT INTO notifications (user_id, type, title, body, ref_id, is_read)
VALUES (?, ?, ?, ?, ?, ?);

-- name: ListNotificationsByUser :many
SELECT * FROM notifications WHERE user_id = ? ORDER BY created_at DESC;

-- name: MarkNotificationAsRead :exec
UPDATE notifications SET is_read = true WHERE id = ?;

-- name: MarkAllAsRead :exec
UPDATE notifications SET is_read = true WHERE user_id = ?;
