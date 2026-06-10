package middleware

import (
	"net/http"
	"strings"

	"sewa-lapangan-voli/config"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware memvalidasi JWT dari header Authorization: Bearer <token>
// dan menyimpan user_id + role ke Gin context.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			config.RespondError(c, http.StatusUnauthorized, "Token tidak ditemukan")
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := config.ParseAccessToken(tokenStr)
		if err != nil {
			config.RespondError(c, http.StatusUnauthorized, "Token tidak valid atau sudah kadaluarsa")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

// RoleGuard membatasi akses endpoint hanya untuk role tertentu.
// Gunakan setelah AuthMiddleware.
//
// Contoh: middleware.RoleGuard("superuser", "owner")
func RoleGuard(allowedRoles ...string) gin.HandlerFunc {
	roleSet := make(map[string]bool, len(allowedRoles))
	for _, r := range allowedRoles {
		roleSet[r] = true
	}

	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			config.RespondError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if !roleSet[role.(string)] {
			config.RespondError(c, http.StatusForbidden, "Akses ditolak")
			c.Abort()
			return
		}

		c.Next()
	}
}
