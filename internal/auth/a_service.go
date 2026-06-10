package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"sewa-lapangan-voli/config"
	"sewa-lapangan-voli/db"

	"golang.org/x/crypto/bcrypt"
)

// ─────────────────────────────────────────
// REGISTER
// ─────────────────────────────────────────

// RegisterService memvalidasi input, membuat user baru, dan
// mengirim email verifikasi (khusus penyewa).
// Untuk owner: status tetap 'pending' sampai upload dokumen & di-approve.
func RegisterService(ctx context.Context, q *db.Queries, req RegisterRequest) (UserResponse, error) {
	// Cek email sudah terdaftar
	existing, err := q.GetUserByEmail(ctx, req.Email)
	if err == nil && existing.ID > 0 {
		return UserResponse{}, errors.New("email sudah terdaftar")
	}

	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse{}, errors.New("gagal memproses password")
	}

	// Buat user — status 'pending' untuk semua role baru
	result, err := q.CreateUser(ctx, db.CreateUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashed),
		Phone:    toNullString(req.Phone),
		Role:     req.Role,
		Status:   "pending",
	})
	if err != nil {
		return UserResponse{}, errors.New("gagal membuat akun")
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return UserResponse{}, errors.New("gagal mendapatkan ID user")
	}

	// Kirim verifikasi email hanya untuk penyewa
	if req.Role == "penyewa" {
		token, err := generateSecureToken()
		if err != nil {
			return UserResponse{}, errors.New("gagal membuat token verifikasi")
		}

		expiresAt := time.Now().Add(24 * time.Hour)
		if err := q.CreateEmailVerification(ctx, db.CreateEmailVerificationParams{
			UserID:    userID,
			Token:     token,
			ExpiresAt: expiresAt,
		}); err != nil {
			return UserResponse{}, errors.New("gagal menyimpan token verifikasi")
		}

		// Kirim email (fire and forget — error tidak menggagalkan register)
		go config.SendVerificationEmail(req.Email, req.Name, token)
	}

	return UserResponse{
		ID:        userID,
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		Role:      req.Role,
		Status:    "pending",
		CreatedAt: time.Now(),
	}, nil
}

// ─────────────────────────────────────────
// VERIFY EMAIL
// ─────────────────────────────────────────

// VerifyEmailService memvalidasi token dan mengaktifkan akun penyewa.
func VerifyEmailService(ctx context.Context, q *db.Queries, req VerifyEmailRequest) error {
	verification, err := q.GetEmailVerificationByToken(ctx, req.Token)
	if err != nil {
		return errors.New("token tidak valid atau sudah kadaluarsa")
	}

	// Tandai token sebagai sudah digunakan
	if err := q.MarkEmailVerificationUsed(ctx, req.Token); err != nil {
		return errors.New("gagal memverifikasi email")
	}

	// Aktifkan user
	if err := q.UpdateUserStatus(ctx, db.UpdateUserStatusParams{
		Status: "active",
		ID:     verification.UserID,
	}); err != nil {
		return errors.New("gagal mengaktifkan akun")
	}

	return nil
}

// ─────────────────────────────────────────
// LOGIN
// ─────────────────────────────────────────

// LoginService memvalidasi kredensial dan mengembalikan access + refresh token.
func LoginService(ctx context.Context, q *db.Queries, req LoginRequest) (LoginResponse, error) {
	user, err := q.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return LoginResponse{}, errors.New("email atau password salah")
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return LoginResponse{}, errors.New("email atau password salah")
	}

	// Cek status akun
	switch user.Status {
	case "pending":
		if user.Role == "penyewa" {
			return LoginResponse{}, errors.New("akun belum diverifikasi, cek email kamu")
		}
		return LoginResponse{}, errors.New("akun kamu sedang menunggu persetujuan admin")
	case "inactive":
		return LoginResponse{}, errors.New("akun kamu telah dinonaktifkan")
	}

	// Generate access token (JWT)
	accessToken, err := config.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return LoginResponse{}, errors.New("gagal membuat access token")
	}

	// Generate refresh token
	refreshToken, err := generateSecureToken()
	if err != nil {
		return LoginResponse{}, errors.New("gagal membuat refresh token")
	}

	// Simpan refresh token ke DB
	expiresAt := time.Now().Add(30 * 24 * time.Hour) // 30 hari
	if err := q.CreateRefreshToken(ctx, db.CreateRefreshTokenParams{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: expiresAt,
	}); err != nil {
		return LoginResponse{}, errors.New("gagal menyimpan sesi login")
	}

	return LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     nullStringToString(user.Phone),
			Role:      user.Role,
			Status:    user.Status,
			CreatedAt: user.CreatedAt,
		},
	}, nil
}

// ─────────────────────────────────────────
// REFRESH TOKEN
// ─────────────────────────────────────────

// RefreshTokenService menukar refresh token lama dengan access token baru.
func RefreshTokenService(ctx context.Context, q *db.Queries, req RefreshTokenRequest) (RefreshResponse, error) {
	rt, err := q.GetRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return RefreshResponse{}, errors.New("refresh token tidak valid atau sudah kadaluarsa")
	}

	user, err := q.GetUserByID(ctx, rt.UserID)
	if err != nil {
		return RefreshResponse{}, errors.New("user tidak ditemukan")
	}

	accessToken, err := config.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return RefreshResponse{}, errors.New("gagal membuat access token")
	}

	return RefreshResponse{AccessToken: accessToken}, nil
}

// ─────────────────────────────────────────
// LOGOUT
// ─────────────────────────────────────────

// LogoutService menghapus refresh token dari DB (invalidate sesi).
func LogoutService(ctx context.Context, q *db.Queries, req LogoutRequest) error {
	if err := q.DeleteRefreshToken(ctx, req.RefreshToken); err != nil {
		return errors.New("gagal logout")
	}
	return nil
}

// ─────────────────────────────────────────
// HELPERS (private)
// ─────────────────────────────────────────

func generateSecureToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func toNullString(s string) db.NullString {
	if s == "" {
		return db.NullString{Valid: false}
	}
	return db.NullString{String: s, Valid: true}
}

func nullStringToString(ns db.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}
