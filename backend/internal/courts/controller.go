package courts

import (
	"net/http"
	"strconv"

	"sewa-lapangan-voli/config"
	sqlc_db "sewa-lapangan-voli/db"

	"github.com/gin-gonic/gin"
)

func CreateCourtController(c *gin.Context) {
	var req CreateCourtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")

	q := sqlc_db.New(config.DB)
	result, err := CreateCourtService(c.Request.Context(), q, userID.(int64), req)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusCreated, "Lapangan berhasil dibuat", result)
}

func GetCourtController(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	q := sqlc_db.New(config.DB)
	result, err := GetCourtService(c.Request.Context(), q, id)
	if err != nil {
		config.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Berhasil", result)
}

func ListCourtsController(c *gin.Context) {
	city := c.Query("city")

	q := sqlc_db.New(config.DB)
	result, err := ListCourtsService(c.Request.Context(), q, city)
	if err != nil {
		config.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Berhasil", result)
}

func ListOwnerCourtsController(c *gin.Context) {
	userID, _ := c.Get("user_id")

	q := sqlc_db.New(config.DB)
	result, err := ListOwnerCourtsService(c.Request.Context(), q, userID.(int64))
	if err != nil {
		config.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Berhasil", result)
}

func UpdateCourtController(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req UpdateCourtRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")
	q := sqlc_db.New(config.DB)
	if err := UpdateCourtService(c.Request.Context(), q, id, userID.(int64), req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Lapangan berhasil diperbarui", nil)
}

func AddCourtImageController(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	var req CourtImageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID, _ := c.Get("user_id")
	q := sqlc_db.New(config.DB)
	if err := AddCourtImageService(c.Request.Context(), q, id, userID.(int64), req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusCreated, "Gambar berhasil ditambahkan", nil)
}

func DeleteCourtController(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, "ID tidak valid")
		return
	}

	userID, _ := c.Get("user_id")
	q := sqlc_db.New(config.DB)
	if err := DeleteCourtService(c.Request.Context(), q, id, userID.(int64)); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Lapangan berhasil dihapus", nil)
}
