-- name: CreateEmailVerification :exec
INSERT INTO email_verifications (user_id, token, expires_at)
VALUES (?, ?, ?);

-- name: GetEmailVerificationByToken :one
SELECT * FROM email_verifications
WHERE token = ?
  AND used_at IS NULL
  AND expires_at > NOW()
LIMIT 1;

-- name: MarkEmailVerificationUsed :exec
UPDATE email_verifications
SET used_at = NOW()
WHERE token = ?;

-- name: DeleteEmailVerificationsByUser :exec
DELETE FROM email_verifications WHERE user_id = ?;
