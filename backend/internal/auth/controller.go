package auth

import (
	"net/http"

	"sewa-lapangan-voli/config"
	sqlc_db "sewa-lapangan-voli/db"

	"github.com/gin-gonic/gin"
)

func RegisterController(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	q := sqlc_db.New(config.DB)
	result, err := RegisterService(c.Request.Context(), q, req)
	if err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusCreated, "Registrasi berhasil", result)
}

func LoginController(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	q := sqlc_db.New(config.DB)
	result, err := LoginService(c.Request.Context(), q, req)
	if err != nil {
		config.RespondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Login berhasil", result)
}

func RefreshTokenController(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	q := sqlc_db.New(config.DB)
	result, err := RefreshTokenService(c.Request.Context(), q, req)
	if err != nil {
		config.RespondError(c, http.StatusUnauthorized, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Token berhasil diperbarui", result)
}

func LogoutController(c *gin.Context) {
	var req LogoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	q := sqlc_db.New(config.DB)
	if err := LogoutService(c.Request.Context(), q, req); err != nil {
		config.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Logout berhasil", nil)
}
