package payments

import (
	"net/http"
	"strconv"

	"sewa-lapangan-voli/config"
	sqlc_db "sewa-lapangan-voli/db"

	"github.com/gin-gonic/gin"
)

func CreatePaymentController(c *gin.Context) {
	var req CreatePaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")

	q := sqlc_db.New(config.DB)
	result, err := CreatePaymentService(c.Request.Context(), q, userID.(int64), req)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusCreated, "Payment berhasil dibuat", result)
}

func MidtransWebhookController(c *gin.Context) {
	var req MidtransWebhookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	q := sqlc_db.New(config.DB)
	if err := ProcessMidtransWebhookService(c.Request.Context(), q, req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Webhook diproses", nil)
}

func GetPaymentController(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	q := sqlc_db.New(config.DB)
	result, err := GetPaymentService(c.Request.Context(), q, id)
	if err != nil {
		config.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Berhasil", result)
}
