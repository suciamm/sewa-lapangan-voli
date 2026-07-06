package notifications

import (
	"context"
	"errors"

	sqlc_db "sewa-lapangan-voli/db"
)

func ListUserNotificationsService(ctx context.Context, q *sqlc_db.Queries, userID int64) ([]NotificationResponse, error) {
	notifs, err := q.ListNotificationsByUser(ctx, userID)
	if err != nil {
		return nil, errors.New("gagal mendapatkan notifikasi")
	}

	resp := make([]NotificationResponse, 0, len(notifs))
	for _, n := range notifs {
		var refID *int64
		if n.RefID.Valid {
			refID = &n.RefID.Int64
		}
		resp = append(resp, NotificationResponse{
			ID:        n.ID,
			UserID:    n.UserID,
			Type:      string(n.Type),
			Title:     n.Title,
			Body:      n.Body,
			RefID:     refID,
			IsRead:    n.IsRead,
			CreatedAt: n.CreatedAt,
		})
	}
	return resp, nil
}

func MarkNotificationReadService(ctx context.Context, q *sqlc_db.Queries, id int64, userID int64) error {
	return q.MarkNotificationAsRead(ctx, id)
}

func MarkAllNotificationsReadService(ctx context.Context, q *sqlc_db.Queries, userID int64) error {
	return q.MarkAllAsRead(ctx, userID)
}
