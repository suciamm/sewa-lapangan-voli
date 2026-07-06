package notifications

import "time"

type NotificationResponse struct {
	ID      int64     `json:"id"`
	UserID  int64     `json:"user_id"`
	Type    string    `json:"type"`
	Title   string    `json:"title"`
	Body    string    `json:"body"`
	RefID   *int64    `json:"ref_id,omitempty"`
	IsRead  bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}
