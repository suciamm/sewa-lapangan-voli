package notifications

import (
	"net/http"
	"strconv"

	"sewa-lapangan-voli/config"
	sqlc_db "sewa-lapangan-voli/db"

	"github.com/gin-gonic/gin"
)

func ListNotificationsController(c *gin.Context) {
	userID, _ := c.Get("user_id")

	q := sqlc_db.New(config.DB)
	result, err := ListUserNotificationsService(c.Request.Context(), q, userID.(int64))
	if err != nil {
		config.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Berhasil", result)
}

func MarkNotificationReadController(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	userID, _ := c.Get("user_id")

	q := sqlc_db.New(config.DB)
	if err := MarkNotificationReadService(c.Request.Context(), q, id, userID.(int64)); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Notifikasi ditandai sebagai dibaca", nil)
}

func MarkAllNotificationsReadController(c *gin.Context) {
	userID, _ := c.Get("user_id")

	q := sqlc_db.New(config.DB)
	if err := MarkAllNotificationsReadService(c.Request.Context(), q, userID.(int64)); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Semua notifikasi ditandai sebagai dibaca", nil)
}
