package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DB adalah koneksi global yang dipakai di seluruh aplikasi.
var DB *sql.DB

// InitDB membuka koneksi ke MySQL dan menjalankan migrasi.
// Dipanggil sekali saat aplikasi start di main.go.
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("[DB] gagal membuka koneksi: %v", err)
	}

	// Connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)

	// Pastikan koneksi benar-benar berhasil
	if err := db.Ping(); err != nil {
		log.Fatalf("[DB] gagal ping database: %v", err)
	}

	log.Println("[DB] koneksi berhasil")

	DB = db

	RunMigrations(db)
}

// RunMigrations menjalankan schema.sql jika tabel belum ada.
func RunMigrations(db *sql.DB) {
	// Cek apakah tabel users sudah ada (untuk menandakan apakah migrasi sudah dijalankan)
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = 'users'").Scan(&count)
	if err != nil {
		log.Fatalf("[MIGRATE] gagal cek tabel: %v", err)
	}

	if count > 0 {
		log.Println("[MIGRATE] tabel sudah ada, skip migrasi")
		return
	}

	// Baca schema.sql
	schemaPath := filepath.Join("..", "schema.sql")
	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		log.Fatalf("[MIGRATE] gagal baca schema.sql: %v", err)
	}

	schema := string(schemaBytes)

	// Pisahkan statement SQL (simple split by ; lalu trim)
	// Hilangkan komentar terlebih dahulu
	reComment := regexp.MustCompile(`--.*$`)
	lines := strings.Split(schema, "\n")
	cleanLines := []string{}
	for _, line := range lines {
		clean := reComment.ReplaceAllString(line, "")
		clean = strings.TrimSpace(clean)
		if clean != "" {
			cleanLines = append(cleanLines, clean)
		}
	}
	cleanSchema := strings.Join(cleanLines, " ")

	// Pisahkan per statement
	statements := strings.Split(cleanSchema, ";")
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		_, err := db.Exec(stmt)
		if err != nil {
			log.Printf("[MIGRATE] peringatan saat eksekusi: %v", err)
			// Tetap lanjutkan, karena beberapa error bisa diabaikan (misal table already exists)
		}
	}

	// Insert default platform setting jika belum ada
	var settingCount int
	err = db.QueryRow("SELECT COUNT(*) FROM platform_settings").Scan(&settingCount)
	if err != nil || settingCount == 0 {
		_, err := db.Exec("INSERT IGNORE INTO platform_settings (fee_percent) VALUES (5.00)")
		if err != nil {
			log.Printf("[MIGRATE] gagal insert setting: %v", err)
		}
	}

	log.Println("[MIGRATE] migrasi selesai")
}
