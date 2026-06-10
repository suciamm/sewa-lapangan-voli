package router

import (
	"sewa-lapangan-voli/internal/auth"
	"sewa-lapangan-voli/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")

	// ─────────────────────────────────────────
	// PUBLIC — tidak perlu JWT
	// ─────────────────────────────────────────
	ApiPublic(api)

	// ─────────────────────────────────────────
	// PROTECTED — semua role yang sudah login
	// ─────────────────────────────────────────
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	ApiProtected(protected)

	// ─────────────────────────────────────────
	// SUPERUSER only
	// ─────────────────────────────────────────
	superuser := api.Group("")
	superuser.Use(middleware.AuthMiddleware(), middleware.RoleGuard("superuser"))
	ApiSuperuser(superuser)

	// ─────────────────────────────────────────
	// OWNER only
	// ─────────────────────────────────────────
	owner := api.Group("")
	owner.Use(middleware.AuthMiddleware(), middleware.RoleGuard("owner"))
	ApiOwner(owner)

	// ─────────────────────────────────────────
	// PENYEWA only
	// ─────────────────────────────────────────
	penyewa := api.Group("")
	penyewa.Use(middleware.AuthMiddleware(), middleware.RoleGuard("penyewa"))
	ApiPenyewa(penyewa)
}

// ─── PUBLIC ──────────────────────────────
func ApiPublic(r *gin.RouterGroup) {
	r.POST("/auth/register", auth.RegisterController)
	r.POST("/auth/verify-email", auth.VerifyEmailController)
	r.POST("/auth/login", auth.LoginController)
	r.POST("/auth/refresh", auth.RefreshTokenController)
}

// ─── PROTECTED (semua role) ───────────────
func ApiProtected(r *gin.RouterGroup) {
	r.POST("/auth/logout", auth.LogoutController)

	// TODO: tambah endpoint lain yang bisa diakses semua role login
}

// ─── SUPERUSER ────────────────────────────
func ApiSuperuser(r *gin.RouterGroup) {
	// TODO: owner verification approve/reject
	// TODO: data pemilik lapangan
	// TODO: report data
}

// ─── OWNER ────────────────────────────────
func ApiOwner(r *gin.RouterGroup) {
	// TODO: upload dokumen verifikasi
	// TODO: CRUD lapangan
	// TODO: report data owner
}

// ─── PENYEWA ──────────────────────────────
func ApiPenyewa(r *gin.RouterGroup) {
	// TODO: cari lapangan
	// TODO: booking
	// TODO: history pesanan
	// TODO: review
}
