package auth

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

	"sewa-lapangan-voli/config"
	sqlc_db "sewa-lapangan-voli/db"

	"golang.org/x/crypto/bcrypt"
)

func RegisterService(ctx context.Context, q *sqlc_db.Queries, req RegisterRequest) (UserResponse, error) {
	existing, err := q.GetUserByEmail(ctx, req.Email)
	if err == nil && existing.ID > 0 {
		return UserResponse{}, errors.New("email sudah terdaftar")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return UserResponse{}, errors.New("gagal memproses password")
	}

	result, err := q.CreateUser(ctx, sqlc_db.CreateUserParams{
		Name:       req.Name,
		Email:      req.Email,
		Password:   string(hashed),
		Phone:      toNullString(req.Phone),
		Role:       sqlc_db.UsersRole(req.Role),
		IsVerified: false,
		Status:     sqlc_db.UsersStatus("active"),
	})
	if err != nil {
		return UserResponse{}, errors.New("gagal membuat akun")
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return UserResponse{}, errors.New("gagal mendapatkan ID user")
	}

	return UserResponse{
		ID:         userID,
		Name:       req.Name,
		Email:      req.Email,
		Phone:      req.Phone,
		Role:       req.Role,
		IsVerified: false,
		Status:     "active",
		CreatedAt:  time.Now(),
	}, nil
}

func LoginService(ctx context.Context, q *sqlc_db.Queries, req LoginRequest) (LoginResponse, error) {
	user, err := q.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return LoginResponse{}, errors.New("email atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return LoginResponse{}, errors.New("email atau password salah")
	}

	if user.Status == "inactive" {
		return LoginResponse{}, errors.New("akun kamu telah dinonaktifkan")
	}

	accessToken, err := config.GenerateAccessToken(user.ID, string(user.Role))
	if err != nil {
		return LoginResponse{}, errors.New("gagal membuat access token")
	}

	refreshToken, err := generateSecureToken()
	if err != nil {
		return LoginResponse{}, errors.New("gagal membuat refresh token")
	}

	expiresAt := time.Now().Add(30 * 24 * time.Hour)
	_, err = q.CreateRefreshToken(ctx, sqlc_db.CreateRefreshTokenParams{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: expiresAt,
	})
	if err != nil {
		return LoginResponse{}, errors.New("gagal menyimpan sesi login")
	}

	return LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: UserResponse{
			ID:         user.ID,
			Name:       user.Name,
			Email:      user.Email,
			Phone:      nullStringToString(user.Phone),
			Role:       string(user.Role),
			IsVerified: user.IsVerified,
			Status:     string(user.Status),
			CreatedAt:  user.CreatedAt,
		},
	}, nil
}

func RefreshTokenService(ctx context.Context, q *sqlc_db.Queries, req RefreshTokenRequest) (RefreshResponse, error) {
	rt, err := q.GetRefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return RefreshResponse{}, errors.New("refresh token tidak valid atau sudah kadaluarsa")
	}

	user, err := q.GetUserByID(ctx, rt.UserID)
	if err != nil {
		return RefreshResponse{}, errors.New("user tidak ditemukan")
	}

	accessToken, err := config.GenerateAccessToken(user.ID, string(user.Role))
	if err != nil {
		return RefreshResponse{}, errors.New("gagal membuat access token")
	}

	return RefreshResponse{AccessToken: accessToken}, nil
}

func LogoutService(ctx context.Context, q *sqlc_db.Queries, req LogoutRequest) error {
	if err := q.DeleteRefreshToken(ctx, req.RefreshToken); err != nil {
		return errors.New("gagal logout")
	}
	return nil
}

func generateSecureToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func toNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func nullStringToString(ns sql.NullString) string {
	if !ns.Valid {
		return ""
	}
	return ns.String
}
