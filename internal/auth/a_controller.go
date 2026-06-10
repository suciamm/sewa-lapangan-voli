package auth

import (
	"net/http"

	"sewa-lapangan-voli/config"
	sqlc_db "sewa-lapangan-voli/db"

	"github.com/gin-gonic/gin"
)

// RegisterController godoc
// @Summary      Register akun baru (owner atau penyewa)
// @Description  Owner: status pending, perlu upload dokumen & approve superuser.
//
//	Penyewa: status pending, perlu verifikasi email.
//
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body RegisterRequest true "Register payload"
// @Success      201  {object} UserResponse
// @Router       /auth/register [post]
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

	msg := "Registrasi berhasil"
	if req.Role == "penyewa" {
		msg = "Registrasi berhasil. Cek email kamu untuk verifikasi akun"
	} else {
		msg = "Registrasi berhasil. Akun kamu sedang menunggu persetujuan admin"
	}

	config.RespondSuccess(c, http.StatusCreated, msg, result)
}

// VerifyEmailController godoc
// @Summary      Verifikasi email penyewa
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body VerifyEmailRequest true "Token verifikasi"
// @Success      200
// @Router       /auth/verify-email [post]
func VerifyEmailController(c *gin.Context) {
	var req VerifyEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	q := sqlc_db.New(config.DB)
	if err := VerifyEmailService(c.Request.Context(), q, req); err != nil {
		config.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	config.RespondSuccess(c, http.StatusOK, "Email berhasil diverifikasi. Silakan login", nil)
}

// LoginController godoc
// @Summary      Login semua role
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body LoginRequest true "Login payload"
// @Success      200  {object} LoginResponse
// @Router       /auth/login [post]
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

// RefreshTokenController godoc
// @Summary      Refresh access token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body RefreshTokenRequest true "Refresh token"
// @Success      200  {object} RefreshResponse
// @Router       /auth/refresh [post]
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

// LogoutController godoc
// @Summary      Logout (invalidate refresh token)
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body body LogoutRequest true "Refresh token yang akan di-invalidate"
// @Success      200
// @Router       /auth/logout [post]
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
