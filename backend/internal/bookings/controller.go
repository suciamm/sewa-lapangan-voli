package bookings

import (
	"net/http"
	"strconv"

	"sewa-lapangan-voli/config"
	sqlc_db "sewa-lapangan-voli/db"

	"github.com/gin-gonic/gin"
)

func CreateBookingController(c *gin.Context) {
	var req CreateBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")

	q := sqlc_db.New(config.DB)
	result, err := CreateBookingService(c.Request.Context(), q, userID.(int64), req)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusCreated, "Booking berhasil dibuat", result)
}

func GetBookingController(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	q := sqlc_db.New(config.DB)
	result, err := GetBookingService(c.Request.Context(), q, id, userID.(int64), role.(string))
	if err != nil {
		config.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Berhasil", result)
}

func ListPenyewaBookingsController(c *gin.Context) {
	userID, _ := c.Get("user_id")

	q := sqlc_db.New(config.DB)
	result, err := ListPenyewaBookingsService(c.Request.Context(), q, userID.(int64))
	if err != nil {
		config.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Berhasil", result)
}

func ListOwnerBookingsController(c *gin.Context) {
	userID, _ := c.Get("user_id")

	q := sqlc_db.New(config.DB)
	result, err := ListOwnerBookingsService(c.Request.Context(), q, userID.(int64))
	if err != nil {
		config.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Berhasil", result)
}

func UpdateBookingStatusController(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req UpdateBookingStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")

	q := sqlc_db.New(config.DB)
	if err := UpdateBookingStatusService(c.Request.Context(), q, id, req.Status, userID.(int64), role.(string)); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Status booking berhasil diperbarui", nil)
}
