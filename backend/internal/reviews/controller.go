package reviews

import (
	"net/http"
	"strconv"

	"sewa-lapangan-voli/config"
	sqlc_db "sewa-lapangan-voli/db"

	"github.com/gin-gonic/gin"
)

func CreateReviewController(c *gin.Context) {
	var req CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")

	q := sqlc_db.New(config.DB)
	result, err := CreateReviewService(c.Request.Context(), q, userID.(int64), req)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusCreated, "Review berhasil dibuat", result)
}

func ListCourtReviewsController(c *gin.Context) {
	courtID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	q := sqlc_db.New(config.DB)
	result, err := ListCourtReviewsService(c.Request.Context(), q, courtID)
	if err != nil {
		config.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Berhasil", result)
}
