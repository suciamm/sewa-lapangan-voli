package router

import (
	"sewa-lapangan-voli/internal/auth"
	"sewa-lapangan-voli/internal/bookings"
	"sewa-lapangan-voli/internal/courts"
	"sewa-lapangan-voli/internal/notifications"
	"sewa-lapangan-voli/internal/payments"
	"sewa-lapangan-voli/internal/reviews"
	"sewa-lapangan-voli/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")

	ApiPublic(api)

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	ApiProtected(protected)

	superuser := api.Group("")
	superuser.Use(middleware.AuthMiddleware(), middleware.RoleGuard("superuser"))
	ApiSuperuser(superuser)

	owner := api.Group("")
	owner.Use(middleware.AuthMiddleware(), middleware.RoleGuard("owner"))
	ApiOwner(owner)

	penyewa := api.Group("")
	penyewa.Use(middleware.AuthMiddleware(), middleware.RoleGuard("penyewa"))
	ApiPenyewa(penyewa)
}

func ApiPublic(r *gin.RouterGroup) {
	r.POST("/auth/register", auth.RegisterController)
	r.POST("/auth/login", auth.LoginController)
	r.POST("/auth/refresh", auth.RefreshTokenController)
	r.POST("/payments/webhook", payments.MidtransWebhookController)
	r.GET("/courts", courts.ListCourtsController)
	r.GET("/courts/:id/reviews", reviews.ListCourtReviewsController)
	r.GET("/courts/:id", courts.GetCourtController)
}

func ApiProtected(r *gin.RouterGroup) {
	r.POST("/auth/logout", auth.LogoutController)
	r.GET("/notifications", notifications.ListNotificationsController)
	r.PUT("/notifications/:id/read", notifications.MarkNotificationReadController)
	r.PUT("/notifications/read-all", notifications.MarkAllNotificationsReadController)
	r.GET("/bookings/:id", bookings.GetBookingController)
	r.PUT("/bookings/:id/status", bookings.UpdateBookingStatusController)
}

func ApiSuperuser(r *gin.RouterGroup) {
	// Add superuser endpoints can be added here
}

func ApiOwner(r *gin.RouterGroup) {
	r.POST("/courts", courts.CreateCourtController)
	r.GET("/my-courts", courts.ListOwnerCourtsController)
	r.PUT("/courts/:id", courts.UpdateCourtController)
	r.DELETE("/courts/:id", courts.DeleteCourtController)
	r.POST("/courts/:id/images", courts.AddCourtImageController)
	r.GET("/owner/bookings", bookings.ListOwnerBookingsController)
}

func ApiPenyewa(r *gin.RouterGroup) {
	r.POST("/bookings", bookings.CreateBookingController)
	r.GET("/penyewa/bookings", bookings.ListPenyewaBookingsController)
	r.POST("/payments", payments.CreatePaymentController)
	r.GET("/payments/:id", payments.GetPaymentController)
	r.POST("/reviews", reviews.CreateReviewController)
}
